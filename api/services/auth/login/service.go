package login

import (
	"github.com/anuragkumar19/connect/api/gen/auth/v1/authv1connect"
	"github.com/rs/zerolog"
)

type Service struct {
	logger *zerolog.Logger
	authv1connect.UnimplementedLoginServiceHandler
}

func New(logger *zerolog.Logger) Service {
	return Service{
		logger: logger,
	}
}
