package services

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis() {
	// Replace with your Redis server address and password
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Default Redis address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})
}

func SetRedis(key string, value any) (redisStatus *redis.StatusCmd, err error) {
	var ctx = context.Background()
	redisStatus = redisClient.Set(ctx, key, value, 0)
	err = redisStatus.Err()
	return
}

func CloseRedis() {
	redisClient.Close()
}
