package smtp

import (
	"github.com/anuragkumar19/connect/env"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

var defaultConfig = &Config{}

func AutoConfig() (*Config, error) {
	cfg := *defaultConfig

	env.Str("SMTP_HOST", &cfg.Host)
	if err := env.Int("SMTP_PORT", &cfg.Port); err != nil {
		return nil, err
	}
	env.Str("SMTP_USERNAME", &cfg.Username)
	env.Str("SMTP_PASSWORD", &cfg.Password)

	return &cfg, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.Host, validation.Required),
		validation.Field(&config.Port, validation.Required),
		validation.Field(&config.Username, validation.Required),
		validation.Field(&config.Password, validation.Required),
	)
}
