package ratelimiterstore

import (
	"context"
	"errors"

	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/pkg/ratelimiter"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type databaseStore struct {
	queries *database.Queries
}

var _ ratelimiter.Store = (*databaseStore)(nil)

func NewDatabaseStore(queries *database.Queries) *databaseStore {
	return &databaseStore{
		queries: queries,
	}
}

func (s *databaseStore) Get(ctx context.Context, id string) (ratelimiter.BucketCtx, error) {
	b, err := s.queries.GetRateLimitBucket(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ratelimiter.BucketCtx{}, ratelimiter.ErrNotFound
		}
		return ratelimiter.BucketCtx{}, err
	}

	return mapDBToBucketCtx(&b), nil
}

func (s *databaseStore) Create(ctx context.Context, bucket ratelimiter.BucketCtx) error {
	return s.queries.CreateRateLimitBucket(ctx, &database.CreateRateLimitBucketParams{
		ID:          bucket.ID,
		LastResetAt: bucket.LastResetAt,
		LastConsumedAt: pgtype.Timestamptz{
			Time:  bucket.LastConsumedAt,
			Valid: bucket.LastConsumedAt.IsZero(),
		},
		Version:  int32(bucket.Revision),
		Consumed: int64(bucket.ConsumedTokenCount),
	})
	// TODO: handle err when primary key conflict
}

func (s *databaseStore) Update(ctx context.Context, bucket ratelimiter.BucketCtx) (ratelimiter.BucketCtx, error) {
	b, err := s.queries.UpdateRateLimitBucket(ctx, &database.UpdateRateLimitBucketParams{
		LastResetAt: bucket.LastResetAt,
		Consumed:    int64(bucket.ConsumedTokenCount),
		LastConsumedAt: pgtype.Timestamptz{
			Time:  bucket.LastConsumedAt,
			Valid: bucket.LastConsumedAt.IsZero(),
		},
		ID:      bucket.ID,
		Version: int32(bucket.Revision),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ratelimiter.BucketCtx{}, ratelimiter.ErrNotFound
		}
		return ratelimiter.BucketCtx{}, err
	}

	return mapDBToBucketCtx(&b), nil
}

func mapDBToBucketCtx(b *database.RateLimitBucket) ratelimiter.BucketCtx {
	return ratelimiter.BucketCtx{
		ID:                 b.ID,
		Revision:           int(b.Version),
		LastResetAt:        b.LastConsumedAt.Time,
		ConsumedTokenCount: int(b.Consumed),
		LastConsumedAt:     b.LastConsumedAt.Time,
	}
}
