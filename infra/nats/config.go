package nats

import "time"

type Config struct {
	Name               string
	Servers            []string
	PingInterval       time.Duration
	MaxPingOutstanding int
}

var defaultConfig = &Config{
	Servers:            []string{},
	Name:               "connect.go.server",
	PingInterval:       time.Minute,
	MaxPingOutstanding: 5,
}

func AutoConfig() *Config {
	cfg := &Config{}

	return cfg
}

func (config Config) Validate() error {
	// TODO:
	return nil
}
