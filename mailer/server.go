package mailer

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
	"gopkg.in/gomail.v2"
	"jaytaylor.com/html2text"

	mailerv1 "github.com/anuragkumar19/connect/events/gen/mailer/v1"
	"github.com/anuragkumar19/connect/infra/nats"
	"github.com/anuragkumar19/connect/infra/smtp"
	"github.com/anuragkumar19/connect/stacktrace"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

var ErrServerNotStarted = errors.New("server already started")
var ErrServerAlreadyStarted = errors.New("server not started")

var (
	mailerJobsStreamName = "mailer:jobs"

	sendEmailSubject      = "mailer.send_email"
	sendEmailConsumerName = "mailer:send_email-consumer"
)

type Server struct {
	mutex             sync.Mutex
	config            *Config
	smtp              *smtp.SMTP
	js                jetstream.JetStream
	sendEmailConsumer jetstream.Consumer

	consumers []jetstream.ConsumeContext
}

func NewServer(ctx context.Context, config *Config, s *smtp.SMTP, nc *nats.NATS) (Server, error) {
	if err := config.Validate(); err != nil {
		return Server{}, fmt.Errorf("invalid mailer config: %w", err)
	}
	js, err := jetstream.New(nc.Conn)
	if err != nil {
		return Server{}, fmt.Errorf("failed to create jetstream client from nats.Conn: %w", err)
	}

	// May cause backward compatibility problems. More investigation needed
	/* From NATS docs:
	Add a stream. Adding a stream is an idempotent function, which means that if a stream does not exist,
	it will be created, and if a stream already exists,
	then the add operation will succeed only if the existing stream matches exactly the attributes specified in the 'add' call.
	*/
	if _, err := js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:        mailerJobsStreamName,
		Description: "Jobs for sending emails",
		Subjects:    []string{sendEmailSubject},
		Retention:   jetstream.WorkQueuePolicy,
		Storage:     jetstream.FileStorage,
		Discard:     jetstream.DiscardNew,
	}); err != nil {
		return Server{}, fmt.Errorf("failed to create or update stream: %w", err)
	}

	// May cause backward compatibility problems. More investigation needed
	/* Docs: Durable consumers as the name implies are meant to last 'forever' and are typically created and deleted administratively
	rather than by the application code which only needs to specify the durable's well known name to use it. */
	sendEmailConsumer, err := js.CreateOrUpdateConsumer(ctx, mailerJobsStreamName, jetstream.ConsumerConfig{
		Name:        sendEmailConsumerName,
		Durable:     sendEmailConsumerName,
		Description: "Consumer to send emails",
		AckWait:     30 * time.Second,
		MaxDeliver:  5,
		BackOff: []time.Duration{
			10 * time.Second,
			30 * time.Second,
			time.Minute,
		},
		FilterSubject: sendEmailSubject,
	})
	if err != nil {
		return Server{}, fmt.Errorf("failed to create or update consumer: %w", err)
	}

	return Server{
		config:            config,
		smtp:              s,
		js:                js,
		sendEmailConsumer: sendEmailConsumer,
		consumers:         []jetstream.ConsumeContext{},
	}, nil
}

func (s *Server) Start() error {
	s.mutex.Lock()
	if len(s.consumers) > 0 {
		return ErrServerAlreadyStarted
	}
	s.mutex.Unlock()
	g := errgroup.Group{}

	log.Info().Int("count", s.config.Concurrency).Msg("staring workers goroutines for mailer")
	for range s.config.Concurrency {
		g.Go(func() error {
			c, err := s.sendEmailConsumer.Consume(s.consumeHandler, jetstream.PullMaxMessages(1))
			if err != nil {
				return fmt.Errorf("failed to register consume handler for %s consumer: %w", s.sendEmailConsumer.CachedInfo().Name, err)
			}
			s.mutex.Lock()
			s.consumers = append(s.consumers, c)
			s.mutex.Unlock()
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	log.Info().Msg("workers goroutines started")
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if len(s.consumers) < 1 {
		return ErrServerNotStarted
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()

	g, ctxNew := errgroup.WithContext(ctx)

	for _, c := range s.consumers {
		g.Go(func() error {
			c.Drain()
			select {
			case <-ctxNew.Done():
				return ctxNew.Err()
			case <-c.Closed():
				return nil
			}
		})
	}
	s.consumers = []jetstream.ConsumeContext{}
	return g.Wait()
}

func (s *Server) consumeHandler(msg jetstream.Msg) {
	/*
		FROM docs:
		After a message is sent from the consumer to a subscribing client application by the server an 'AckWait' timer is started.
		This timer is deleted when either a positive (Ack()) or a termination (Term()) acknowledgement is received from the client application.
		The timer gets reset upon reception of an in-progress (inProgress()) acknowledgement.
	*/
	// TODO: use msg.InProgress() for controlling timeouts and ackWait reset

	defer func() {
		if v := recover(); v != nil {
			log.Error().RawJSON("stack", stacktrace.Marshal(debug.Stack())).Err(fmt.Errorf("mailer consume handler panicked: %v", v)).Msg("panicked recovered")
		}
	}()
	// TODO: use context with timeout = AckWait
	meta, err := msg.Metadata()
	if err != nil {
		if err := msg.TermWithReason("failed to retrieve metadata from message"); err != nil {
			log.Error().Err(err).Msg("failed to terminate msg")
			return
		}
		log.Error().Err(err).Msg("failed to retrieve metadata from nats message, message terminated")
		return
	}

	l := log.With().Uint64("stream_sequence", meta.Sequence.Stream).Uint64("consumer_sequence", meta.Sequence.Consumer).Str("domain", meta.Domain).Logger()

	l.Info().Msg("message received")

	var sendEmailMsg mailerv1.SendEmail
	if err := proto.Unmarshal(msg.Data(), &sendEmailMsg); err != nil {
		if err := msg.TermWithReason(fmt.Sprintf("failed to parse data into expected format: %s", err.Error())); err != nil {
			log.Error().Err(err).Msg("failed to terminate msg")
			return
		}
		l.Error().Err(err).Msg("failed to parse data into expected format, message terminated")
		return
	}

	l.Info().Msg("sending email")
	if err := s.sendMail(&sendEmailMsg); err != nil {
		if err := msg.Nak(); err != nil {
			l.Error().Err(err).Msg("failed to nak message")
			return
		}
		l.Error().Err(err).Msg("failed to send email, message naked")
		return
	}
	if err := msg.DoubleAck(context.Background()); err != nil {
		if err := msg.Nak(); err != nil {
			l.Error().Err(err).Msg("failed to nak message")
			return
		}
		l.Error().Err(err).Msg("failed to ack message, message naked")
		return
	}
	l.Info().Msg("email sent successfully")
}

func (s *Server) sendMail(msg *mailerv1.SendEmail) error {
	// TODO: different "from" field of email based on email type, attachment support
	mailMsg := gomail.NewMessage()

	mailMsg.SetAddressHeader("From", s.config.FromAddress, s.config.FromDisplayName)
	mailMsg.SetHeader("To", msg.To)
	if s.config.ReplyTo != "" {
		mailMsg.SetHeader("Reply-To", s.config.ReplyTo)
	}
	mailMsg.SetHeader("Subject", msg.Subject)
	mailMsg.SetDateHeader("Date", time.Now())

	plainBody, err := html2text.FromString(msg.GetHtml(), html2text.Options{
		PrettyTables: true,
	})
	if err != nil {
		return err
	}

	mailMsg.SetBody("text/plain", plainBody)
	mailMsg.AddAlternative("text/html", msg.GetHtml())
	return s.smtp.DialAndSend(mailMsg)
}
