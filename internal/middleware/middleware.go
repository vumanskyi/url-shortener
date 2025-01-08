package middleware

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type rateLimiter struct {
	redisClient *redis.Client
	maxRequests int           // Maximum requests allowed
	expiredAt   time.Duration // Time window (e.g., 1 minute)
}

// NewRateLimiter creates and returns a new rateLimiter instance.
// The rateLimiter uses a Redis client to store request counts and enforces a limit
// on the number of requests allowed within a specified time window.
func NewRateLimiter(redisClient *redis.Client, maxRequests int, expiredAt time.Duration) *rateLimiter {
	return &rateLimiter{
		redisClient,
		maxRequests,
		expiredAt,
	}
}

func (rl *rateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // You can modify this to use an API key or other unique identifier

		// Generate Redis key based on IP
		key := "rate_limit:" + ip

		// Get current request count from Redis
		count, err := rl.redisClient.Get(context.Background(), key).Int()
		if err != nil && err != redis.Nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// If the count exceeds the limit, reject the request
		if count >= rl.maxRequests {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Increment the request count in Redis
		_, err = rl.redisClient.Incr(context.Background(), key).Result()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set an expiration on the key (after the window duration)
		if count == 0 {
			_, err := rl.redisClient.Expire(context.Background(), key, rl.expiredAt).Result()
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

		// Proceed with the request if within rate limit
		next.ServeHTTP(w, r)
	})
}
