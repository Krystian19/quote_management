package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RDS struct {
	client *redis.Client
	ctx    context.Context
}

func NewRDS(ctx context.Context, addr string) (*RDS, error) {
	conn := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := conn.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RDS{client: conn, ctx: ctx}, nil
}

func (r *RDS) GetClient() *redis.Client {
	return r.client
}

// SetValueInCache : Persists a key-value on the cache
func (r *RDS) SetValueInCache(key string, value interface{}, expiration time.Duration) error {
	if err := r.client.Set(r.ctx, key, value, expiration).Err(); err != nil {
		return fmt.Errorf("redis: failed to persist value %s", err)
	}

	return nil
}

// GetValueFromCache : Gets the value from cache that corresponds to the provided key
func (r *RDS) GetValueFromCache(key string) (*string, error) {
	value, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	return &value, nil
}

// RemoveFromCache : Removes a value from cache
func (r *RDS) RemoveFromCache(key string) error {
	if err := r.client.Del(r.ctx, key).Err(); err != nil {
		return fmt.Errorf("redis: failed to remove value from cache %s", err)
	}

	return nil
}

// IsValueInCache : Checks if a value is present on the cache
func (r *RDS) IsValueInCache(key string) (bool, error) {
	value, err := r.GetValueFromCache(key)
	return value != nil, err
}
