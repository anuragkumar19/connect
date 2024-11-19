package server

import (
	"strconv"
	"time"

	"github.com/anuragkumar19/connect/pkg/env"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Host string
	Port string

	HandlerTimeout time.Duration
}

var defaultConfig = &Config{
	Host:           "0.0.0.0",
	Port:           "5005",
	HandlerTimeout: 30 * time.Second,
}

func AutoConfig() (*Config, error) {
	cfg := *defaultConfig

	env.Str("SERVER_HOST", &cfg.Host)
	env.Str("NATS_PORT", &cfg.Port)
	if err := env.Duration("SERVER_HANDLER_TIMEOUT", &cfg.HandlerTimeout); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.Host, validation.Required),
		validation.Field(&config.Port, validation.Required, validation.NewStringRule(func(v string) bool {
			p, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if p < 0 || p > 65535 {
				return false
			}

			return true
		}, "invalid port")),
		validation.Field(&config.HandlerTimeout, validation.Required, validation.Min(time.Millisecond)),
	)
}
