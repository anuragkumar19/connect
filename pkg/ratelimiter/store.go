package ratelimiter

import (
	"context"
)

type Store interface {
	Get(ctx context.Context, id string) (BucketCtx, error)
	Update(context.Context, BucketCtx) error
}

type StoreTX interface {
	BeginFunc(context.Context, func(Store) error) error
}
