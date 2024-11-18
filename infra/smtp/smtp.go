package smtp

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"
)

type SMTP struct {
	*gomail.Dialer
}

func New(ctx context.Context, config *Config) (SMTP, error) {
	if err := config.Validate(); err != nil {
		return SMTP{}, fmt.Errorf("smtp config validation failed: %w", err)
	}

	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	sc, err := dialer.Dial()
	if err != nil {
		return SMTP{}, fmt.Errorf("failed to authenticate with smtp server: %w", err)
	}
	if err := sc.Close(); err != nil {
		return SMTP{}, fmt.Errorf("failed to close smtp conn: %w", err)
	}

	log.Info().Msg("smtp auth successful")

	return SMTP{
		Dialer: dialer,
	}, nil
}
