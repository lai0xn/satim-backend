package store

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string) (string, error)
	Publish(ctx context.Context, requestID string, payload interface{}) error
	Subscribe(ctx context.Context, requestID string) *redis.PubSub
}

func NewRedisStore() Store {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisStore{
		client: client,
	}
}

type RedisStore struct {
	client *redis.Client
}

func (s *RedisStore) Set(ctx context.Context, key string, value any) error {
	err := s.client.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisStore) Get(ctx context.Context, key string) (string, error) {
	result, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *RedisStore) Publish(ctx context.Context, requestID string, payload interface{}) error {
	err := s.client.Publish(ctx, requestID, payload).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisStore) Subscribe(ctx context.Context, requestID string) *redis.PubSub {
	subscriber := s.client.Subscribe(ctx, requestID)
	return subscriber
}
