package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

type NATS struct {
	*nats.Conn

	closeCh chan struct{}
}

func New(ctx context.Context, config *Config) (NATS, error) {
	if err := config.Validate(); err != nil {
		return NATS{}, fmt.Errorf("invalid nats config: %w", err)
	}
	closeCh := make(chan struct{}, 1)
	// TODO: Add context support manually ref: https://docs.nats.io/using-nats/developer/tutorials/custom_dialer
	nc, err := nats.Connect(
		config.Servers,
		nats.PingInterval(config.PingInterval),
		nats.MaxPingsOutstanding(config.MaxPingOutstanding),
		nats.Name(config.Name),
		nats.Token(config.Token),
		nats.ClosedHandler(func(c *nats.Conn) {
			closeCh <- struct{}{}
		}),
	)
	if err != nil {
		return NATS{}, fmt.Errorf("failed to connect to nats servers: %w", err)
	}

	log.Info().Msg("connected to nats servers")

	return NATS{
		Conn:    nc,
		closeCh: closeCh,
	}, nil
}

func (n *NATS) DrainAndWaitForClose(ctx context.Context) error {
	if err := n.Conn.Drain(); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-n.closeCh:
		return nil
	}
}
