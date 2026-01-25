package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/xerdin442/api-practice/internal/config"
)

type Cache struct {
	Redis *redis.Client
}

func New(c *config.Config) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		Password: c.RedisPassword,
	})

	return &Cache{Redis: client}
}

func (c *Cache) SetJTI(ctx context.Context, key, value string, exp time.Time) error {
	return c.Redis.Set(ctx, key, value, time.Until(exp)).Err()
}

func (c *Cache) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	n, err := c.Redis.Exists(ctx, token).Result()
	if err != nil {
		return false, err
	}

	return n > 0, nil
}
