package ratelimiter

import (
	"context"

	"github.com/anuragkumar19/connect/pkg/ratelimiter"
)

type Service interface {
	UserTriggeredEmailBucket(ctx context.Context, email string) (ratelimiter.Bucket[string], error)
}
