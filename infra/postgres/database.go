package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Postgres struct {
	*pgxpool.Pool
}

func New(ctx context.Context, config *Config) (Postgres, error) {
	if err := config.Validate(); err != nil {
		return Postgres{}, fmt.Errorf("invalid postgres config: %w", err)
	}
	dbConf, err := pgxpool.ParseConfig(config.URI)
	if err != nil {
		return Postgres{}, fmt.Errorf("cannot parse postgres config: %w", err)
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
		return Postgres{}, fmt.Errorf("failed to create postgres pool: %w", err)
	}

	// ping db
	if err := dbPool.Ping(ctx); err != nil {
		return Postgres{}, fmt.Errorf("failed to ping postgres: %w", err)
	}

	log.Info().Msg("connected to postgres")

	return Postgres{Pool: dbPool}, nil
}
