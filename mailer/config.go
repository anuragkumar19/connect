package mailer

import (
	"runtime"

	"github.com/anuragkumar19/connect/env"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	Concurrency     int
	FromAddress     string
	FromDisplayName string
	ReplyTo         string
}

var defaultConfig = &Config{
	Concurrency:     runtime.NumCPU(),
	FromAddress:     "",
	FromDisplayName: "",
	ReplyTo:         "",
}

func AutoConfig() (*Config, error) {
	cfg := *defaultConfig

	if err := env.Int("MAILER_CONCURRENCY", &cfg.Concurrency); err != nil {
		return nil, err
	}
	env.Str("MAILER_FROM_ADDRESS", &cfg.FromAddress)
	env.Str("MAILER_FROM_DISPLAY_NAME", &cfg.FromDisplayName)
	env.Str("MAILER_REPLY_TO", &cfg.ReplyTo)
	return &cfg, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.Concurrency, validation.Required, validation.Min(1)),
		validation.Field(&config.FromAddress, validation.Required, is.EmailFormat),
		validation.Field(&config.FromDisplayName, validation.Required),
		validation.Field(&config.ReplyTo, validation.Required, is.EmailFormat),
	)
}
