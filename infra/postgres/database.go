package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Postgres struct {
	*pgxpool.Pool
}

func New(ctx context.Context, config *Config, logger *zerolog.Logger) (Postgres, error) {
	if err := config.Validate(); err != nil {
		return Postgres{}, errors.Join(errors.New("invalid postgres config"), err)
	}
	dbConf, err := pgxpool.ParseConfig(config.URI)
	if err != nil {
		return Postgres{}, errors.Join(errors.New("cannot parse postgres config"), err)
	}
	dbConf.MaxConnLifetime = config.MaxConnLifetime
	dbConf.MaxConnLifetimeJitter = config.MaxConnLifetimeJitter
	dbConf.MaxConnIdleTime = config.MaxConnIdleTime
	dbConf.MaxConns = config.MaxConns
	dbConf.MinConns = config.MinConns
	dbConf.HealthCheckPeriod = config.HealthCheckPeriod

	// create pool
	dbPool, err := pgxpool.NewWithConfig(ctx, dbConf)
	if err != nil {
		return Postgres{}, errors.Join(errors.New("failed to create postgres pool"), err)
	}

	// ping db
	if err := dbPool.Ping(ctx); err != nil {
		return Postgres{}, errors.Join(errors.New("failed to ping postgres"), err)
	}

	return Postgres{Pool: dbPool}, nil
}
