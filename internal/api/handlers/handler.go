package handlers

import (
	"github.com/hibiken/asynq"
	"github.com/xerdin442/api-practice/internal/cache"
	"github.com/xerdin442/api-practice/internal/service"
)

type RouteHandler struct {
	services   *service.Manager
	cache      *cache.Redis
	tasksQueue *asynq.Client
}

func New(svc *service.Manager, c *cache.Redis, q *asynq.Client) *RouteHandler {
	return &RouteHandler{services: svc, cache: c, tasksQueue: q}
}
