package handlers

import (
	"github.com/hibiken/asynq"
	"github.com/xerdin442/api-practice/internal/cache"
	"github.com/xerdin442/api-practice/internal/service"
)

type RouteHandler struct {
	services   *service.Manager
	cache      *cache.Cache
	tasksQueue *asynq.Client
}

func New(svc *service.Manager, c *cache.Cache, q *asynq.Client) *RouteHandler {
	return &RouteHandler{services: svc, cache: c, tasksQueue: q}
}
