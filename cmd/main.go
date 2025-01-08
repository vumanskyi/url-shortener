package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"vumanskyi/url-shortener/config"
	"vumanskyi/url-shortener/internal/handler"
	"vumanskyi/url-shortener/internal/middleware"
	"vumanskyi/url-shortener/pkg/redis"
)

func main() {
	appConfig := config.NewAppConfig()
	redisClient := redis.InitClient(appConfig.RedisConfig)

	router := chi.NewRouter()
	rateLimiter := middleware.NewRateLimiter(redisClient, appConfig.RateLimit.MaxRequest, appConfig.RateLimit.ExpiredAt)
	router.Use(rateLimiter.Limit)

	h := handler.NewHandler(redisClient)
	router.Get("/{shortUrl}", h.GetShortenedUrl)
	router.Post("/", h.GetShortenedUrl)

	slog.Info("Start application.")

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port), router); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
