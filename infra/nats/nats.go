package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type NATS struct {
	*nats.Conn
}

func New(ctx context.Context, config *Config, logger *zerolog.Logger) (NATS, error) {
	if err := config.Validate(); err != nil {
		return NATS{}, fmt.Errorf("invalid nats config: %w", err)
	}
	// TODO: Add context support manually
	nc, err := nats.Connect(
		config.Servers,
		nats.PingInterval(config.PingInterval),
		nats.MaxPingsOutstanding(config.MaxPingOutstanding),
		nats.Name(config.Name),
		nats.Token(config.Token),
	)
	if err != nil {
		return NATS{}, fmt.Errorf("failed to connect to nats servers: %w", err)
	}

	logger.Info().Msg("connected to nats servers")

	return NATS{
		Conn: nc,
	}, nil
}
