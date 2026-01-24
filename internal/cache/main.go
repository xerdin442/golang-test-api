package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/xerdin442/api-practice/internal/config"
)

type Redis struct {
	Client *redis.Client
	cfg    *config.Config
}

func NewRedis(addr, password string, c *config.Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	return &Redis{Client: client, cfg: c}
}

func (r *Redis) SetJTI(ctx context.Context, key, value string, exp time.Time) error {
	return r.Client.Set(ctx, key, value, time.Until(exp)).Err()
}

func (r *Redis) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	n, err := r.Client.Exists(ctx, token).Result()
	if err != nil {
		return false, err
	}

	return n > 0, nil
}
