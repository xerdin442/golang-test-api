package service

import (
	"github.com/xerdin442/api-practice/internal/config"
	repo "github.com/xerdin442/api-practice/internal/repository"
)

var secrets = config.Load()

type Manager struct {
	Event *EventService
	User  *UserService
}

func NewManager(r *repo.Registry) *Manager {
	return &Manager{
		Event: NewEventService(r.Event),
		User:  NewUserService(r.User),
	}
}
