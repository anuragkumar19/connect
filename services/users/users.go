package users

import (
	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/mailer"
	"github.com/anuragkumar19/connect/services/ratelimiter"
)

type Users struct {
	store       database.Store
	mailer      *mailer.Client
	rateLimiter ratelimiter.Service
}

var _ Service = (*Users)(nil)

func New(store database.Store, m *mailer.Client, r ratelimiter.Service) Users {
	return Users{
		store:       store,
		mailer:      m,
		rateLimiter: r,
	}
}
