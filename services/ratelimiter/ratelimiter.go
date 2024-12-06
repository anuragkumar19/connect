package ratelimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/pkg/ratelimiter"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Ratelimiter struct {
	store *databaseStoreTx

	userTriggeredEmailLimiter *ratelimiter.BasicLimiter[string]
}

var _ Service = (*Ratelimiter)(nil)

func New(queries *database.Queries, db *pgxpool.Pool) (Ratelimiter, error) {
	store := newDatabaseStore(queries, db)

	userTriggeredEmailLimiter, err := ratelimiter.NewBasicLimiter(&ratelimiter.BasicLimiterOption[string]{
		HashFunc:   ratelimiter.String,
		Label:      "user_triggered_email",
		Limit:      20,
		ResetAfter: 24 * time.Hour,
		Store:      store,
		BackOffs:   []time.Duration{time.Minute, 2 * time.Minute, 3 * time.Minute, 5 * time.Minute, 10 * time.Minute, 30 * time.Minute},
	})
	if err != nil {
		return Ratelimiter{}, fmt.Errorf("failed to create useTriggeredEmailLimiter: %w", err)
	}

	return Ratelimiter{
		store:                     store,
		userTriggeredEmailLimiter: userTriggeredEmailLimiter,
	}, nil
}

func (s *Ratelimiter) UserTriggeredEmailBucket(ctx context.Context, email string) *ratelimiter.Bucket[string] {
	return s.userTriggeredEmailLimiter.Bucket(ctx, email)
}
