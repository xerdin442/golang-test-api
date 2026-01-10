package service

import repo "github.com/xerdin442/api-practice/internal/repository"

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
