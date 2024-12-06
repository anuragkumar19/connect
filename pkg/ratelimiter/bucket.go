package ratelimiter

import (
	"context"
	"time"
)

type BucketCtx struct {
	ID                 string    `json:"id"`
	LastResetAt        time.Time `json:"last_reset_at"`
	ConsumedTokenCount int       `json:"consumed_token_count"`
	LastConsumedAt     time.Time `json:"last_consumed_at"`
}

func NewBucketCtx(id string) BucketCtx {
	return BucketCtx{
		ID:                 id,
		LastResetAt:        time.Now(),
		ConsumedTokenCount: 0,
		LastConsumedAt:     time.Time{},
	}
}

type Bucket[T any] struct {
	l  *BasicLimiter[T]
	id string
}

func (b *Bucket[T]) Consume(ctx context.Context) error {
	now := time.Now()

	return b.l.store.BeginFunc(ctx, func(s Store) error {
		bucket, err := s.Get(ctx, b.id)
		if err != nil {
			return err
		}

		if bucket.LastResetAt.Add(b.l.resetAfter).Before(now) {
			bucket.ConsumedTokenCount = 0
			bucket.LastResetAt = now
			bucket.LastConsumedAt = time.Time{}
		}

		if bucket.ConsumedTokenCount >= b.l.limit {
			tryAfter := bucket.LastResetAt.Add(b.l.resetAfter)
			return &RateLimitError{
				Remaining: 0,
				ResetAt:   tryAfter,
				TryAfter:  tryAfter,
			}
		}

		backoffLen := len(b.l.backOffs)
		if backoffLen > 0 && bucket.ConsumedTokenCount > 0 {
			i := bucket.ConsumedTokenCount - 1
			if bucket.ConsumedTokenCount > backoffLen {
				i = backoffLen - 1
			}
			d := b.l.backOffs[i]

			tryAfter := bucket.LastConsumedAt.Add(d)
			resetAt := bucket.LastResetAt.Add(b.l.resetAfter)
			if resetAt.Before(tryAfter) {
				tryAfter = resetAt
			}
			if tryAfter.After(now) {
				return &RateLimitError{
					Remaining: b.l.limit - bucket.ConsumedTokenCount,
					ResetAt:   resetAt,
					TryAfter:  tryAfter,
				}
			}
		}

		bucket.ConsumedTokenCount++
		bucket.LastConsumedAt = now

		if err := s.Update(ctx, bucket); err != nil {
			return err
		}
		return nil
	})
}
