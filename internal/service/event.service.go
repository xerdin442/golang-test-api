package service

import repo "github.com/xerdin442/api-practice/internal/repository"

type EventService struct {
	repo repo.EventRepoInterface
}

func NewEventService(r repo.EventRepoInterface) *EventService {
	return &EventService{repo: r}
}
