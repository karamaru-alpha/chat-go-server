package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("redis:%s", os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})
}
