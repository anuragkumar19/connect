package consumer

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/anuragkumar19/connect/pkg/stacktrace"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

var (
	ErrServerNotStarted     = errors.New("server already started")
	ErrServerAlreadyStarted = errors.New("server not started")
)

type Consumer[T proto.Message] struct {
	concurrency  int
	consumer     jetstream.Consumer
	consumeFunc  ConsumeFunc[T]
	logger       *zerolog.Logger
	pool         sync.Pool
	timeout      time.Duration
	mu           sync.Mutex
	consumerCtxs []jetstream.ConsumeContext
	baseCtx      context.Context
}

type ConsumeFunc[T proto.Message] func(context.Context, T) error

func New[T proto.Message](name string, concurrency int, consumer jetstream.Consumer, consumeFunc ConsumeFunc[T], newFn func() T, timeout time.Duration) (Consumer[T], error) {
	if name == "" {
		return Consumer[T]{}, errors.New("name cannot be empty")
	}

	if concurrency < 1 {
		return Consumer[T]{}, errors.New("concurrency must be greater than or equal to 1")
	}

	if consumer == nil {
		return Consumer[T]{}, errors.New("jetstream consumer cannot be nil")
	}

	if consumeFunc == nil {
		return Consumer[T]{}, errors.New("consumeFunc cannot be nil")
	}

	if newFn == nil {
		return Consumer[T]{}, errors.New("newFn cannot be nil")
	}

	if timeout < 0 {
		return Consumer[T]{}, errors.New("timeout cannot be negative")
	}

	logger := log.With().Str("consumer_name", name).Logger()

	return Consumer[T]{
		concurrency: concurrency,
		logger:      &logger,
		consumer:    consumer,
		consumeFunc: consumeFunc,
		pool: sync.Pool{
			New: func() any {
				return newFn()
			},
		},
		timeout:      timeout,
		mu:           sync.Mutex{},
		consumerCtxs: make([]jetstream.ConsumeContext, 0, concurrency),
		baseCtx:      nil,
	}, nil
}

func (c *Consumer[T]) Start(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.consumerCtxs) > 0 {
		return ErrServerAlreadyStarted
	}

	g := errgroup.Group{}

	c.logger.Info().Int("count", c.concurrency).Msg("staring workers goroutines")
	for range c.concurrency {
		g.Go(func() error {
			consumerCtx, err := c.consumer.Consume(c.handler, jetstream.PullMaxMessages(1))
			if err != nil {
				return fmt.Errorf("failed to register consume handler for %s consumer: %w", c.consumer.CachedInfo().Name, err)
			}
			c.consumerCtxs = append(c.consumerCtxs, consumerCtx)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	c.baseCtx = ctx
	c.logger.Info().Msg("workers goroutines started")
	return nil
}

func (c *Consumer[T]) Shutdown(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.consumerCtxs) < 1 {
		return ErrServerNotStarted
	}

	g, ctxNew := errgroup.WithContext(ctx)

	for _, consumerCtx := range c.consumerCtxs {
		g.Go(func() error {
			consumerCtx.Drain()
			select {
			case <-ctxNew.Done():
				return ctxNew.Err()
			case <-consumerCtx.Closed():
				return nil
			}
		})
	}
	c.consumerCtxs = make([]jetstream.ConsumeContext, 0, c.concurrency)
	return g.Wait()
}

func (c *Consumer[T]) handler(msg jetstream.Msg) {
	defer func() {
		if v := recover(); v != nil {
			c.logger.Error().RawJSON("stack", stacktrace.Marshal(debug.Stack())).Err(fmt.Errorf("consume handler panicked: %v", v)).Msg("panicked recovered")
		}
	}()

	ctx, cancel := context.WithTimeout(c.baseCtx, c.timeout)
	defer cancel()

	meta, err := msg.Metadata()
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to retrieve metadata from nats message")
		if err := msg.TermWithReason("failed to retrieve metadata from message"); err != nil {
			c.logger.Error().Err(err).Msg("failed to terminate msg")
		}
		return
	}

	l := c.logger.With().Uint64("stream_sequence", meta.Sequence.Stream).Uint64("consumer_sequence", meta.Sequence.Consumer).Logger()

	v := c.pool.Get().(T)
	if err := proto.Unmarshal(msg.Data(), v); err != nil {
		l.Error().Err(err).Msg("failed to parse data into expected format")
		if err := msg.TermWithReason(fmt.Sprintf("failed to parse data into expected format: %s", err.Error())); err != nil {
			l.Error().Err(err).Msg("failed to terminate msg")
		}
		return
	}

	if err := c.consumeFunc(ctx, v); err != nil {
		l.Error().Err(err).Msg("failed to consume message")
		if err := msg.Nak(); err != nil {
			l.Error().Err(err).Msg("failed to nak message")
		}
		return
	}

	if err := msg.DoubleAck(ctx); err != nil {
		l.Error().Err(err).Msg("failed to ack message")
		if err := msg.TermWithReason(fmt.Sprintf("failed to ack message: %s", err.Error())); err != nil {
			l.Error().Err(err).Msg("failed to terminate msg")
		}
		return
	}
}
