package service

import (
	"crypto/sha256"
	"net/url"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encodeBase62(num int64) string {
	result := ""
	for num > 0 {
		remainder := num % 62
		result = string(base62Chars[remainder]) + result
		num = num / 62
	}
	return result
}

// GenerateShortURL takes a URL and generates a short URL in the form of a Base62 encoded string
// from the first 8 characters of the SHA256 hash of the original URL.
func GenerateShortURL(originalURL string) string {
	if !IsValidURL(originalURL) {
		return ""
	}

	hash := sha256.Sum256([]byte(originalURL))

	// Convert the hash to Base62
	num := int64(hash[0])<<56 | int64(hash[1])<<48 | int64(hash[2])<<40 | int64(hash[3])<<32 |
		int64(hash[4])<<24 | int64(hash[5])<<16 | int64(hash[6])<<8 | int64(hash[7])

	return encodeBase62(num)[:8] // Return first 8 characters
}

// IsValidURL takes a raw URL and checks if it is valid or not.
// A URL is valid if it can be parsed, has a scheme and a host.
func IsValidURL(rawURL string) bool {
	if rawURL == "" {
		return false
	}

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}
