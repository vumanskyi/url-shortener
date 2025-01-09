package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"vumanskyi/url-shortener/internal/service"
)

func Test_handler_GetShortenedUrl(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	h := NewHandler(rdb)

	tests := []struct {
		name           string
		shortUrl       string
		mockRedisFunc  func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "valid short URL",
			shortUrl: "abcd123",
			mockRedisFunc: func() {
				mock.ExpectGet("abcd123").SetVal("https://example.com")
			},
			expectedStatus: http.StatusFound,
			expectedBody:   "",
		},
		{
			name:     "short URL not found",
			shortUrl: "notfound",
			mockRedisFunc: func() {
				mock.ExpectGet("notfound").RedisNil()
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Short URL not found\n",
		},
		{
			name:     "invalid URL stored in Redis",
			shortUrl: "invalid",
			mockRedisFunc: func() {
				mock.ExpectGet("invalid").SetVal("invalid-url")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Invalid URL stored in Redis\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRedisFunc()

			req := httptest.NewRequest("GET", "/{shortUrl}", nil)
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("shortUrl", tt.shortUrl)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			rec := httptest.NewRecorder()
			h.GetShortenedUrl(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
			if tt.expectedBody != "" {
				body := rec.Body.String()
				assert.Equal(t, tt.expectedBody, body)
			}
		})
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func Test_handler_CreateShortenedUrl(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	h := NewHandler(rdb)

	validURL := "https://example.com"
	shortURL := service.GenerateShortURL(validURL)
	expiration := time.Second * 3600

	tests := []struct {
		name           string
		requestBody    shortenRequest
		mockRedisFunc  func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "valid request",
			requestBody: shortenRequest{
				Url:       validURL,
				ExpiredAt: expiration,
			},
			mockRedisFunc: func() {
				mock.ExpectSet(shortURL, validURL, expiration).SetVal("OK")
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "invalid URL",
			requestBody: shortenRequest{
				Url: "invalid-url",
			},
			mockRedisFunc:  func() {},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "Invalid URL\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			tt.mockRedisFunc()
			h.CreateShortenedUrl(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
			if tt.expectedBody != "" {
				body := rec.Body.String()
				assert.Equal(t, tt.expectedBody, body)
			}
		})
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
