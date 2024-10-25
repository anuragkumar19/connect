package smtp

import (
	"errors"

	"github.com/rs/zerolog"
	"gopkg.in/gomail.v2"
)

type SMTP struct {
	*gomail.Dialer
}

func NewDialer(config *Config, logger *zerolog.Logger) (SMTP, error) {
	if err := config.Validate(); err != nil {
		return SMTP{}, errors.Join(errors.New("smtp config validation failed"), err)
	}

	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	if _, err := dialer.Dial(); err != nil {
		return SMTP{}, errors.Join(errors.New("failed to authenticate with smtp server"), err)
	}

	return SMTP{
		Dialer: dialer,
	}, nil
}
