package ratelimiter

import (
	"github.com/anuragkumar19/connect/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Ratelimiter struct {
	store *databaseStoreTx
}

func New(queries *database.Queries, db *pgxpool.Pool) (Ratelimiter, error) {
	return Ratelimiter{
		store: newDatabaseStore(queries, db),
	}, nil
}
