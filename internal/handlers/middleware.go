package handlers

import (
	"net/http"
	"strings"

	"rate-limiter/internal/metrics"
	"rate-limiter/internal/ratelimiter"
)

func RateLimitMiddleware(limiter ratelimiter.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP := strings.Split(r.RemoteAddr, ":")[0]
			allowed, err := limiter.Allow(clientIP)

			metrics.TotalRequests.Inc()

			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if !allowed {
				metrics.RateLimitedRequests.Inc()
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome! You're not rate-limited."))
}
