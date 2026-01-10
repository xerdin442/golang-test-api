package handlers

import (
	"github.com/xerdin442/api-practice/internal/cache"
	"github.com/xerdin442/api-practice/internal/service"
)

type RouteHandler struct {
	services *service.Manager
	cache    *cache.Redis
}

func New(svc *service.Manager, c *cache.Redis) *RouteHandler {
	return &RouteHandler{services: svc, cache: c}
}
