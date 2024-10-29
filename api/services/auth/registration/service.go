package registration

import (
	"github.com/anuragkumar19/connect/api/gen/auth/v1/authv1connect"
	"github.com/anuragkumar19/connect/database"
	"github.com/rs/zerolog"
)

type Service struct {
	logger *zerolog.Logger
	store  *database.Queries
}

var _ authv1connect.RegistrationServiceHandler = (*Service)(nil)

func New(logger *zerolog.Logger) Service {
	return Service{
		logger: logger,
	}
}
