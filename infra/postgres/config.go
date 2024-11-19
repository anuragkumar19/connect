package postgres

import (
	"runtime"
	"time"

	"github.com/anuragkumar19/connect/pkg/env"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	URI string

	MaxConnLifetime       time.Duration
	MaxConnLifetimeJitter time.Duration
	MaxConnIdleTime       time.Duration
	MaxConns              int32
	MinConns              int32
	HealthCheckPeriod     time.Duration
}

var defaultConfig = &Config{
	URI:                   "",
	MaxConnLifetime:       time.Hour,
	MaxConnLifetimeJitter: 5 * time.Minute,
	MaxConnIdleTime:       30 * time.Minute,
	MaxConns:              int32(runtime.NumCPU()),
	MinConns:              4,
	HealthCheckPeriod:     time.Minute,
}

func AutoConfig() (*Config, error) {
	cfg := *defaultConfig

	env.Str("POSTGRES_URI", &cfg.URI)
	if err := env.Duration("POSTGRES_MAX_CONN_LIFETIME", &cfg.MaxConnLifetime); err != nil {
		return nil, err
	}
	if err := env.Duration("POSTGRES_MAX_CONN_LIFETIME_JITTER", &cfg.MaxConnLifetimeJitter); err != nil {
		return nil, err
	}
	if err := env.Duration("POSTGRES_MAX_CONN_IDLE_TIME", &cfg.MaxConnIdleTime); err != nil {
		return nil, err
	}
	if err := env.Int32("POSTGRES_MAX_CONNS", &cfg.MaxConns); err != nil {
		return nil, err
	}
	if err := env.Int32("POSTGRES_MIX_CONNS", &cfg.MinConns); err != nil {
		return nil, err
	}
	if err := env.Duration("POSTGRES_HEALTH_CHECK_PERIOD", &cfg.HealthCheckPeriod); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.URI, validation.Required, is.URL),
		validation.Field(&config.MaxConnLifetime, validation.Required, validation.Min(time.Minute)),
		validation.Field(&config.MaxConnLifetimeJitter, validation.Required, validation.Min(time.Second)),
		validation.Field(&config.MaxConnIdleTime, validation.Required, validation.Min(time.Minute)),
		validation.Field(&config.MinConns, validation.Required),
		validation.Field(&config.MaxConns, validation.Required, validation.Min(config.MaxConns).Error("maxConn should be greater that minConn")),
		validation.Field(&config.HealthCheckPeriod, validation.Required, validation.Min(time.Second)),
	)
}
