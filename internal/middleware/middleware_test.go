package middleware

import (
	"fmt"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_rateLimiter_Limit(t *testing.T) {
	ip := "192.0.2.1:1234"
	// Create a mock Redis client
	rdb, mock := redismock.NewClientMock()

	// Define the time window and request limit
	maxRequests := 5
	expiredAt := time.Minute

	// Create the rate limiter instance
	rl := NewRateLimiter(rdb, maxRequests, expiredAt)

	// Create a dummy handler to pass into the middleware
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Create a new HTTP request and response recorder
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	// Mock Redis responses for a successful request
	mock.ExpectGet(fmt.Sprintf("rate_limit:%s", ip)).SetVal("0")
	mock.ExpectIncr(fmt.Sprintf("rate_limit:%s", ip)).SetVal(1)
	mock.ExpectExpire(fmt.Sprintf("rate_limit:%s", ip), expiredAt).SetVal(true)

	// Call the rate limiter middleware
	rl.Limit(nextHandler).ServeHTTP(rec, req)

	// Verify the response
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Equal(t, "OK", rec.Body.String())

	// Ensure all mock expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
