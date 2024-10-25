package postgres

import (
	"runtime"
	"time"
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

func AutoConfig() *Config {
	cfg := &Config{}

	return cfg
}

func (config Config) Validate() error {
	// TODO:
	return nil
}
