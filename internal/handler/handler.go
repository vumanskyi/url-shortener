package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"vumanskyi/url-shortener/internal/service"
)

type handler struct {
}

type shortenRequest struct {
	Url string `json:"url"`
}

type shortenResponse struct {
	ShortUrl string `json:"short_url"`
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) GetShortenedUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")

	originalURL := service.LookupOriginalURL(shortUrl)

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func (h *handler) CreateShortenedUrl(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	json.NewDecoder(r.Body).Decode(&req)

	slog.Info("Save", req)

	res := shortenResponse{
		ShortUrl: service.GenerateShortURL(req.Url),
	}

	json.NewEncoder(w).Encode(res)
}
