package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type dbtx interface {
	DBTX
	Begin(ctx context.Context) (pgx.Tx, error)
}

type Store interface {
	Querier
	BeginFunc(ctx context.Context, fn func(Store) error) (err error)
}

type store struct {
	*Queries
	db dbtx
}

func NewStore(db dbtx) Store {
	return &store{
		db:      db,
		Queries: New(db),
	}
}

func (s *store) BeginFunc(ctx context.Context, fn func(Store) error) (err error) {
	return pgx.BeginFunc(ctx, s.db, func(tx pgx.Tx) error {
		return fn(NewStore(tx))
	})
}
