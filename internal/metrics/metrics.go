package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_requests",
			Help: "Total HTTP requests received",
		},
	)

	RateLimitedRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "rate_limited_requests",
			Help: "Number of requests rate-limited",
		},
	)
)

func Register() {
	prometheus.MustRegister(TotalRequests)
	prometheus.MustRegister(RateLimitedRequests)
}
