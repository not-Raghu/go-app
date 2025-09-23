package db

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func InitRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "",
		DB:         0,
		MaxRetries: 3,
	})

	_, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		log.Fatal("error with redis client")
	}

	return RedisClient
}
