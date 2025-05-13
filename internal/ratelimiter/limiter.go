package ratelimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	Allow(clientID string) (bool, error)
}

type RedisStore struct {
	client *redis.Client
	rate   int
	window time.Duration
}

func NewRedisStore(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func NewRateLimiter(client *redis.Client, rate int, window time.Duration) *RedisStore {
	return &RedisStore{
		client: client,
		rate:   rate,
		window: window,
	}
}

func (r *RedisStore) Allow(clientID string) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("rate_limit:%s:%d", clientID, time.Now().Unix()/int64(r.window.Seconds()))

	count, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		r.client.Expire(ctx, key, r.window)
	}

	return count <= int64(r.rate), nil
}
