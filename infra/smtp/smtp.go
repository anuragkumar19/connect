package smtp

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"gopkg.in/gomail.v2"
)

type SMTP struct {
	*gomail.Dialer
}

func New(ctx context.Context, config *Config, logger *zerolog.Logger) (SMTP, error) {
	if err := config.Validate(); err != nil {
		return SMTP{}, fmt.Errorf("smtp config validation failed: %w", err)
	}

	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	if _, err := dialer.Dial(); err != nil {
		return SMTP{}, fmt.Errorf("failed to authenticate with smtp server: %w", err)
	}

	logger.Info().Msg("smtp auth successful")

	return SMTP{
		Dialer: dialer,
	}, nil
}
