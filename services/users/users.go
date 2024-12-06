package users

import (
	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/mailer"
	"github.com/anuragkumar19/connect/services/ratelimiter"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users struct {
	db          *pgxpool.Pool
	queries     *database.Queries
	mailer      *mailer.Client
	rateLimiter ratelimiter.Service
}

var _ Service = (*Users)(nil)

func New(db *pgxpool.Pool, queries *database.Queries, m *mailer.Client, r ratelimiter.Service) Users {
	return Users{
		db:          db,
		queries:     queries,
		mailer:      m,
		rateLimiter: r,
	}
}
