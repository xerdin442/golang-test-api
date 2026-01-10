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

func (s *EventService) UpdateEvent(ctx context.Context, dto dto.UpdateEventRequest, eventID, userID int32) (database.Event, error) {
	if dto.Datetime.Before(time.Now()) {
		return database.Event{}, ErrInvalidDate
	}

	event, _ := s.repo.GetEvent(ctx, eventID)
	if event.OwnerID != userID {
		return database.Event{}, ErrOwnerRestrictedAction
	}

	arg := database.UpdateEventParams{
		Name:        dto.Name,
		Description: dto.Description,
		Location:    dto.Location,
		Datetime:    dto.Datetime,
		OwnerID:     userID,
		ID:          eventID,
	}

	result, err := s.repo.UpdateEvent(ctx, arg)
	if err != nil {
		return database.Event{}, err
	}

	updatedEventID, _ := result.LastInsertId()
	return s.repo.GetEvent(ctx, int32(updatedEventID))
}

func (s *EventService) ListEvents(ctx context.Context) ([]database.Event, error) {
	result, err := s.repo.ListEvents(ctx)
	if err != nil {
		return []database.Event{}, err
	}

	return result, nil
}

func (s *EventService) GetEvent(ctx context.Context, eventID int32) (database.Event, error) {
	result, err := s.repo.GetEvent(ctx, eventID)
	if err != nil {
		return database.Event{}, err
	}

	return result, nil
}

func (s *EventService) DeleteEvent(ctx context.Context, eventID, userID int32) error {
	event, _ := s.repo.GetEvent(ctx, eventID)
	if event.OwnerID != userID {
		return ErrOwnerRestrictedAction
	}

	if err := s.repo.DeleteEvent(ctx, eventID); err != nil {
		return err
	}

	return nil
}

func (s *EventService) ReserveTicket(ctx context.Context, userID, eventID int32) error {
	args := database.AddAttendeeParams{
		UserID:  userID,
		EventID: int32(eventID),
	}

	_, err := s.repo.AddAttendee(ctx, args)
	if err != nil {
		return err
	}

	return nil
}

func (s *EventService) RevokeTicket(ctx context.Context, userID, eventID int32) error {
	args := database.RemoveAttendeeParams{
		UserID:  userID,
		EventID: int32(eventID),
	}

	if err := s.repo.RemoveAttendee(ctx, args); err != nil {
		return err
	}

	return nil
}

func (s *EventService) GetEventAttendees(ctx context.Context, userID, eventID int32) ([]database.GetEventAttendeesRow, error) {
	event, _ := s.repo.GetEvent(ctx, eventID)
	if event.OwnerID != userID {
		return []database.GetEventAttendeesRow{}, ErrOwnerRestrictedAction
	}

	return s.repo.GetEventAttendees(ctx, eventID)
}
