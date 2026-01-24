package service

import (
	"github.com/xerdin442/api-practice/internal/config"
	repo "github.com/xerdin442/api-practice/internal/repository"
)

type Manager struct {
	Event *EventService
	User  *UserService
}

func NewManager(r *repo.Registry, cfg *config.Config) *Manager {
	return &Manager{
		Event: NewEventService(r.Event),
		User:  NewUserService(r.User, cfg),
	}
}
