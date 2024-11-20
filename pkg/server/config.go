package server

import (
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	PortValidationRule = validation.NewStringRule(func(v string) bool {
		p, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if p < 0 || p > 65535 {
			return false
		}

		return true
	}, "invalid port")

	HostValidationRule = validation.NewStringRule(func(v string) bool {
		// TODO: impl
		return true
	}, "invalid host")
)

type Config struct {
	Name string
	Host string
	Port string

	HandlerTimeout time.Duration
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.Name, validation.Required),
		validation.Field(&config.Host, validation.Required, HostValidationRule),
		validation.Field(&config.Port, validation.Required, PortValidationRule),
		validation.Field(&config.HandlerTimeout, validation.Required, validation.Min(time.Millisecond)),
	)
}
