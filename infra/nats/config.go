package nats

import (
	"time"

	"github.com/anuragkumar19/connect/env"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Name               string
	Servers            string
	PingInterval       time.Duration
	MaxPingOutstanding int
	Token              string
}

var defaultConfig = &Config{
	Name:               "connect.go.server",
	Servers:            "",
	PingInterval:       time.Minute,
	MaxPingOutstanding: 5,
	Token:              "",
}

func AutoConfig() (*Config, error) {
	cfg := *defaultConfig

	env.Str("NATS_NAME", &cfg.Name)
	env.Str("NATS_SERVERS", &cfg.Servers)
	if err := env.Duration("NATS_PING_INTERVAL", &cfg.PingInterval); err != nil {
		return nil, err
	}
	if err := env.Int("NATS_MAX_PING_OUTSTANDING", &cfg.MaxPingOutstanding); err != nil {
		return nil, err
	}
	env.Str("NATS_TOKEN", &cfg.Token)
	return &cfg, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.Name, validation.Required),
		validation.Field(&config.Servers, validation.Required, validation.NewStringRule(func(v string) bool {
			// s := strings.Split(v, ",")
			// for _, ss := range s {
			// 	if err := is.URL.Validate(ss); err != nil {
			// 		return false
			// 	}
			// }
			return true
		}, "invalid servers url string")),
		validation.Field(&config.PingInterval, validation.Required, validation.Min(time.Second)),
		validation.Field(&config.MaxPingOutstanding, validation.Required),
		validation.Field(&config.Token),
	)
}
