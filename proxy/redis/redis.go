package redis

import "github.com/redis/go-redis/v9"

var Client redis.Client

func NewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Client = *client
}
