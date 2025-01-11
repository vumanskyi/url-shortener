package config

import (
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
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
	env := os.Getenv("APP_ENV")
	if env != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		slog.Error("Failed to parse REDIS_DB environment variable", "error", err.Error())
		redisDB = 0
	}

	redisConfig := RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	}

	rateLimit := RateLimit{
		MaxRequest: 15,
		ExpiredAt:  1 * time.Minute,
	}

	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		appPort = "8080"
	}

	return &AppConfig{
		Port:        appPort,
		RedisConfig: redisConfig,
		RateLimit:   rateLimit,
	}
}
