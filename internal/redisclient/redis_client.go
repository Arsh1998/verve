package redisclient

import (
	"context"
	"fmt"
	"time"
	"verve/config"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

// Initialize connects to the Redis server
func Initialize(config config.Config) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	// Test the connection
	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

// IsDuplicate checks if the ID already exists in Redis
func IsDuplicate(id string) bool {
	exists, err := Client.SetNX(Ctx, id, true, 1*time.Minute).Result()
	if err != nil {
		panic("Redis operation failed: " + err.Error())
	}
	return !exists
}
