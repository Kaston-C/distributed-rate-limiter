package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"rate-limiter/internal/handlers"
	"rate-limiter/internal/metrics"
	"rate-limiter/internal/ratelimiter"
)

func main() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rateLimit := 10 // requests per minute
	store := ratelimiter.NewRedisStore(redisAddr)
	limiter := ratelimiter.NewRateLimiter(store, rateLimit, time.Minute)
	metrics.Register()

	http.Handle("/", handlers.RateLimitMiddleware(limiter)(http.HandlerFunc(handlers.RootHandler)))
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
