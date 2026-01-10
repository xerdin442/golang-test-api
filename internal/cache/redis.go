package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/xerdin442/api-practice/internal/env"
)

type Redis struct{ Client *redis.Client }

func NewRedis() *Redis {
	addr := env.GetStr("REDIS_ADDR")
	password := env.GetStr("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{Addr: addr, Password: password})
	return &Redis{Client: rdb}
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
