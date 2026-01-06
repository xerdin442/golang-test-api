package service

import (
	"context"
	"time"

	database "github.com/xerdin442/api-practice/internal/adapters/generated"
	"github.com/xerdin442/api-practice/internal/api/dto"

	repo "github.com/xerdin442/api-practice/internal/repository"
)

type EventService struct {
	repo repo.EventRepoInterface
}

func NewEventService(r repo.EventRepoInterface) *EventService {
	return &EventService{repo: r}
}

func (s *EventService) CreateEvent(ctx context.Context, dto dto.CreateEventRequest, userID int32) (database.Event, error) {
	if dto.Datetime.Before(time.Now()) {
		return database.Event{}, ErrInvalidDate
	}

	arg := database.CreateEventParams{
		Name:        dto.Name,
		Description: dto.Description,
		Location:    dto.Location,
		Datetime:    dto.Datetime,
		OwnerID:     userID,
	}

	result, err := s.repo.CreateEvent(ctx, arg)
	if err != nil {
		return database.Event{}, err
	}

	eventID, _ := result.LastInsertId()
	return s.repo.GetEvent(ctx, int32(eventID))
}
