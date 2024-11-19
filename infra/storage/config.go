package storage

import (
	"github.com/anuragkumar19/connect/pkg/env"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	Endpoint  string
	AccessKey string
	Secret    string
	UseSSL    bool
	Region    string
}

var defaultConfig = &Config{
	Endpoint:  "",
	AccessKey: "",
	Secret:    "",
	UseSSL:    false,
	Region:    "auto",
}

func AutoConfig() (*Config, error) {
	cfg := *defaultConfig

	env.Str("STORAGE_ENDPOINT", &cfg.Endpoint)
	env.Str("STORAGE_ACCESS_KEY", &cfg.AccessKey)
	env.Str("STORAGE_SECRET", &cfg.Secret)
	if err := env.Bool("STORAGE_USE_SSL", &cfg.UseSSL); err != nil {
		return nil, err
	}
	env.Str("STORAGE_REGION", &cfg.Region)

	return &cfg, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.Endpoint, validation.Required, is.URL),
		validation.Field(&config.AccessKey, validation.Required),
		validation.Field(&config.Secret, validation.Required),
		validation.Field(&config.UseSSL),
		validation.Field(&config.Region, validation.Required),
	)
}
