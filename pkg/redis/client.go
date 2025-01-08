package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"vumanskyi/url-shortener/config"
)

// InitClient returns a Redis client to interact with the Redis server.
// The client is configured from the given 	.
func InitClient(redisConfig config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	return rdb
}
