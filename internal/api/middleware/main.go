package middleware

import (
	"github.com/xerdin442/api-practice/internal/cache"
	"github.com/xerdin442/api-practice/internal/config"
)

type Middleware struct {
	cfg   *config.Config
	cache *cache.Cache
}

func New(c *config.Config, r *cache.Cache) *Middleware {
	return &Middleware{cfg: c, cache: r}
}
