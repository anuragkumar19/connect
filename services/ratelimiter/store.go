package ratelimiter

import (
	"context"
	"errors"

	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/pkg/ratelimiter"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type databaseStoreTx struct {
	queries *database.Queries
	db      *pgxpool.Pool
}

var _ ratelimiter.StoreTX = (*databaseStoreTx)(nil)

func newDatabaseStore(queries *database.Queries, db *pgxpool.Pool) *databaseStoreTx {
	return &databaseStoreTx{
		queries: queries,
		db:      db,
	}
}

func (s *databaseStoreTx) BeginFunc(ctx context.Context, txFunc func(ratelimiter.Store) error) error {
	return pgx.BeginFunc(ctx, s.db, func(tx pgx.Tx) error {
		return txFunc(&databaseStore{
			queries: s.queries.WithTx(tx),
		})
	})
}

type databaseStore struct {
	queries *database.Queries
}

var _ ratelimiter.Store = (*databaseStore)(nil)

func (s *databaseStore) Get(ctx context.Context, id string) (ratelimiter.BucketCtx, error) {
	bucket, err := s.queries.GetRateLimitBucketForUpdate(ctx, id)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return ratelimiter.BucketCtx{}, err
		}
		bCtx := ratelimiter.NewBucketCtx(id)
		newBucket, err := s.queries.CreateRateLimitBucket(ctx, &database.CreateRateLimitBucketParams{
			ID:          bCtx.ID,
			LastResetAt: bCtx.LastResetAt,
			Consumed:    int64(bCtx.ConsumedTokenCount),
			LastConsumedAt: pgtype.Timestamptz{
				Time:  bCtx.LastConsumedAt,
				Valid: bCtx.LastResetAt.IsZero(),
			},
		})
		if err != nil {
			return ratelimiter.BucketCtx{}, err
		}
		bucket = newBucket
	}

	return ratelimiter.BucketCtx{
		ID:                 bucket.ID,
		LastResetAt:        bucket.LastResetAt,
		ConsumedTokenCount: int(bucket.Consumed),
		LastConsumedAt:     bucket.LastConsumedAt.Time,
	}, nil
}

func (s *databaseStore) Update(ctx context.Context, bucket ratelimiter.BucketCtx) error {
	return s.queries.UpdateRateLimitBucket(ctx, &database.UpdateRateLimitBucketParams{
		LastResetAt: bucket.LastResetAt,
		Consumed:    int64(bucket.ConsumedTokenCount),
		LastConsumedAt: pgtype.Timestamptz{
			Time:  bucket.LastConsumedAt,
			Valid: bucket.LastConsumedAt.IsZero(),
		},
		ID: bucket.ID,
	})
}
