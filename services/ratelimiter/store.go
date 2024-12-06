package ratelimiter

import (
	"context"
	"errors"

	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/pkg/ratelimiter"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type databaseStoreTx struct {
	store database.Store
}

var _ ratelimiter.StoreTX = (*databaseStoreTx)(nil)

func newDatabaseStore(store database.Store) *databaseStoreTx {
	return &databaseStoreTx{
		store: store,
	}
}

func (s *databaseStoreTx) BeginFunc(ctx context.Context, txFunc func(ratelimiter.Store) error) error {
	return s.store.BeginFunc(ctx, func(store database.Store) error {
		return txFunc(&databaseStore{
			store: store,
		})
	})
}

type databaseStore struct {
	store database.Store
}

var _ ratelimiter.Store = (*databaseStore)(nil)

func (s *databaseStore) Get(ctx context.Context, id string) (ratelimiter.BucketCtx, error) {
	bucket, err := s.store.GetRateLimitBucketForUpdate(ctx, id)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return ratelimiter.BucketCtx{}, err
		}
		bCtx := ratelimiter.NewBucketCtx(id)
		newBucket, err := s.store.CreateRateLimitBucket(ctx, &database.CreateRateLimitBucketParams{
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
	return s.store.UpdateRateLimitBucket(ctx, &database.UpdateRateLimitBucketParams{
		LastResetAt: bucket.LastResetAt,
		Consumed:    int64(bucket.ConsumedTokenCount),
		LastConsumedAt: pgtype.Timestamptz{
			Time:  bucket.LastConsumedAt,
			Valid: bucket.LastConsumedAt.IsZero(),
		},
		ID: bucket.ID,
	})
}
