package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/xerdin442/api-practice/internal/config"
)

var secrets = config.Load()

type Redis struct{ Client *redis.Client }

func NewRedis() *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     secrets.RedisAddr,
		Password: secrets.RedisPassword,
	})
	return &Redis{Client: client}
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
