package ratelimiter

import (
	"context"
	"errors"
	"time"
)

type HashFunc[T any] func(T) string

type BasicLimiter[T any] struct {
	label      string
	hashFunc   HashFunc[T]
	resetAfter time.Duration
	backOffs   []time.Duration
	store      StoreTX
	limit      int
}

type BasicLimiterOption[T any] struct {
	Label      string
	Limit      int
	ResetAfter time.Duration
	HashFunc   HashFunc[T]
	Store      StoreTX
	BackOffs   []time.Duration
}

func NewBasicLimiter[T any](option *BasicLimiterOption[T]) (*BasicLimiter[T], error) {
	if option.Label == "" {
		return nil, errors.New("label cannot be empty")
	}
	if option.Limit < 0 {
		return nil, errors.New("limit cannot be negative")
	}
	if option.ResetAfter <= 0 {
		return nil, errors.New("resetAfter cannot be zero or negative")
	}
	if option.HashFunc == nil {
		return nil, errors.New("hashFunc cannot be nil")
	}
	if len(option.BackOffs) >= option.Limit {
		return nil, errors.New("backOffs length must be less than limit")
	}
	if option.Store == nil {
		return nil, errors.New("store cannot be nil")
	}

	return &BasicLimiter[T]{
		label:      option.Label,
		hashFunc:   option.HashFunc,
		resetAfter: option.ResetAfter,
		backOffs:   option.BackOffs,
		store:      option.Store,
		limit:      option.Limit,
	}, nil
}

func (l *BasicLimiter[T]) Bucket(ctx context.Context, v T) *Bucket[T] {
	id := l.label + ":" + l.hashFunc(v)
	return &Bucket[T]{
		l:  l,
		id: id,
	}
}

func String(s string) string {
	return s
}
