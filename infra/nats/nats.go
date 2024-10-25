package nats

import (
	"context"
	"errors"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type NATS struct {
	*nats.Conn
}

func New(ctx context.Context, config *Config, logger *zerolog.Logger) (NATS, error) {
	if err := config.Validate(); err != nil {
		return NATS{}, errors.Join(errors.New("invalid nats config"), err)
	}
	// TODO: Add context support manually
	nc, err := nats.Connect(
		strings.Join(config.Servers, ","),
		nats.PingInterval(config.PingInterval),
		nats.MaxPingsOutstanding(config.MaxPingOutstanding),
		nats.Name(config.Name),
	)
	if err != nil {
		return NATS{}, errors.Join(errors.New("failed to connect to nats servers"), err)
	}

	return NATS{
		Conn: nc,
	}, nil
}
