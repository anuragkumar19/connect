package users

import (
	"github.com/anuragkumar19/connect/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users struct {
	db      *pgxpool.Pool
	queries *database.Queries
}

var _ Service = (*Users)(nil)

func New(db *pgxpool.Pool, queries *database.Queries) Users {
	return Users{
		db:      db,
		queries: queries,
	}
}
