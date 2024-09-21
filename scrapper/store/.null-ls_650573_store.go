package store

import "context"

type Store interface {
	Set(ctx context.Context, key string, value any)
	Get(ctx context.Context, key string)
}

func NewRedisStore() Store {
}
