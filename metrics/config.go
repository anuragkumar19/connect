package metrics

import (
	"time"

	"github.com/anuragkumar19/connect/pkg/env"
	"github.com/anuragkumar19/connect/pkg/server"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Host string
	Port string

	HandlerTimeout time.Duration
}

var defaultConfig = &Config{
	Host:           "0.0.0.0",
	Port:           "5006",
	HandlerTimeout: 30 * time.Second,
}

func AutoConfig() (*Config, error) {
	cfg := *defaultConfig

	env.Str("METRICS_SERVER_HOST", &cfg.Host)
	env.Str("METRICS_SERVER_PORT", &cfg.Port)
	if err := env.Duration("METRICS_SERVER_HANDLER_TIMEOUT", &cfg.HandlerTimeout); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.Host, validation.Required, server.HostValidationRule),
		validation.Field(&config.Port, validation.Required, server.PortValidationRule),
		validation.Field(&config.HandlerTimeout, validation.Required, validation.Min(time.Millisecond)),
	)
}
