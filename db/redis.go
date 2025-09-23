package db

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func InitRedis() *redis.Client {

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	maxtriesdb, _ := strconv.Atoi(os.Getenv("REDIS_MAX_TRIES"))

	RedisClient = redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PASS"),
		DB:         db,
		MaxRetries: maxtriesdb,
	})

	_, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		log.Fatal("error with redis client")
	}

	return RedisClient
}
