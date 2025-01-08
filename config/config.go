package config

import "time"

type AppConfig struct {
	Host        string
	Port        string
	RedisConfig RedisConfig
	RateLimit   RateLimit
}

type RateLimit struct {
	MaxRequest int
	ExpiredAt  time.Duration
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// NewAppConfig initializes and returns a pointer to an AppConfig
func NewAppConfig() *AppConfig {
	redisConfig := RedisConfig{
		Host:     "localhost",
		Port:     "6379",
		Password: "",
		DB:       0,
	}

	rateLimit := RateLimit{
		MaxRequest: 15,
		ExpiredAt:  1 * time.Minute,
	}

	return &AppConfig{
		Host:        "localhost",
		Port:        "3000",
		RedisConfig: redisConfig,
		RateLimit:   rateLimit,
	}
}
