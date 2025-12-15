package handlers

import "github.com/xerdin442/api-practice/internal/service"

type RouteHandler struct {
	services *service.Manager
}

func New(svc *service.Manager) *RouteHandler {
	return &RouteHandler{services: svc}
}
