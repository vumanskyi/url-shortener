package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"time"
	"vumanskyi/url-shortener/internal/service"
)

const prefix = "sht"

type handler struct {
	redisClient *redis.Client
}

type shortenRequest struct {
	Url       string        `json:"url"`
	ExpiredAt time.Duration `json:"expired_at"`
}

type shortenResponse struct {
	ShortUrl string `json:"short_url"`
}

func NewHandler(redisClient *redis.Client) *handler {
	return &handler{redisClient}
}

func (h *handler) GetShortenedUrl(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL parameter from the request
	shortUrl := chi.URLParam(r, "shortUrl")

	// Get the original URL from Redis
	originalURL, err := h.redisClient.Get(r.Context(), fmt.Sprintf("%s:%s", prefix, shortUrl)).Result()
	if err != nil {
		if err == redis.Nil {
			slog.Error("Short URL not found", "error", err.Error())
			http.Error(w, "Short URL not found", http.StatusNotFound)
		} else {
			slog.Error("Server error", "error", err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Validate the original URL before redirecting
	if !service.IsValidURL(originalURL) {
		http.Error(w, "Invalid URL stored in Redis", http.StatusInternalServerError)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func (h *handler) CreateShortenedUrl(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest

	// Decode the JSON request body and handle errors
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Invalid request payload", "error", err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if !service.IsValidURL(req.Url) {
		slog.Error("Invalid URL", "URL", req.Url)
		http.Error(w, "Invalid URL", http.StatusUnprocessableEntity)
		return
	}

	expiredAt := req.ExpiredAt

	// Generate the short URL
	shortUrl := service.GenerateShortURL(req.Url)
	slog.Info("Input request", "URL", req.Url, "ExpiredAt", expiredAt, "shortURL", shortUrl)

	if err := h.redisClient.Set(r.Context(), fmt.Sprintf("%s:%s", prefix, shortUrl), req.Url, expiredAt).Err(); err != nil {
		slog.Error("Failed to save to Redis", "error", err.Error())
		http.Error(w, "Failed to save to Redis", http.StatusInternalServerError)
		return
	}

	// Prepare and send the response
	res := shortenResponse{ShortUrl: shortUrl}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.Error("Failed to encode response", "error", err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
