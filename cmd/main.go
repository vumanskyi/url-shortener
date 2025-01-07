package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"vumanskyi/url-shortener/config"
	"vumanskyi/url-shortener/internal/handler"
)

const port = ":3000"

func main() {
	appConfig := config.AppConfig{}
	router := chi.NewRouter()

	h := handler.NewHandler()

	router.Get("/{shortUrl}", h.GetShortenedUrl)
	router.Post("/", h.GetShortenedUrl)

	slog.Info("Start application.")

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port), router); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
