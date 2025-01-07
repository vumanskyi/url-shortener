package service

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortURL(originalURL string) string {
	hash := sha256.Sum256([]byte(originalURL))
	return base64.URLEncoding.EncodeToString(hash[:6])
}

func LookupOriginalURL(url string) string {
	// Placeholder logic - integrate with Redis/PostgreSQL
	return "https://example.com"
}
