package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	database "github.com/xerdin442/api-practice/internal/adapters/generated"
	"github.com/xerdin442/api-practice/internal/api/dto"
	"github.com/xerdin442/api-practice/internal/util"
)

func TestCreateEvent(t *testing.T) {
	t.Run("Invalid Date", func(t *testing.T) {
		createEventDto := dto.CreateEventRequest{
			Datetime: time.Now().AddDate(-1, 2, 3),
		}

		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		_, err := svc.CreateEvent(context.Background(), createEventDto, int32(1))

		assert.ErrorIs(t, err, util.ErrInvalidDate)
	})

	t.Run("Success", func(t *testing.T) {
		createEventDto := dto.CreateEventRequest{
			Name:        "Test Event",
			Description: "This is a test event",
			Location:    "Test Location",
			Datetime:    time.Now().Add(time.Hour),
		}

		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("CreateEvent", mock.Anything, mock.Anything).
			Return(new(mockDbResult), nil)

		mockRepo.On("GetEvent", mock.Anything, mock.Anything).
			Return(database.Event{}, nil)

		_, err := svc.CreateEvent(context.Background(), createEventDto, int32(1))

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateEvent(t *testing.T) {
	updateEventDto := dto.UpdateEventRequest{
		Name:     "Updated Event Name",
		Datetime: time.Now().Add(time.Hour),
	}

	t.Run("Invalid Date", func(t *testing.T) {
		updateEventDto := dto.UpdateEventRequest{
			Datetime: time.Now().AddDate(-1, 2, 3),
		}

		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		_, err := svc.UpdateEvent(context.Background(), updateEventDto, int32(1), int32(1))

		assert.ErrorIs(t, err, util.ErrInvalidDate)
	})

	t.Run("Event Not Found", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("GetEvent", mock.Anything, mock.Anything).
			Return(database.Event{}, sql.ErrNoRows)

		_, err := svc.UpdateEvent(context.Background(), updateEventDto, int32(1), int32(1))

		assert.ErrorIs(t, err, util.ErrEventNotFound)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Unauthorized Access", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("GetEvent", mock.Anything, mock.Anything).
			Return(database.Event{OwnerID: 2}, nil)

		_, err := svc.UpdateEvent(context.Background(), updateEventDto, int32(1), int32(1))

		assert.ErrorIs(t, err, util.ErrOwnerRestrictedAction)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("UpdateEvent", mock.Anything, mock.Anything).
			Return(new(mockDbResult), nil)

		mockRepo.On("GetEvent", mock.Anything, mock.Anything).
			Return(database.Event{OwnerID: 1}, nil)

		_, err := svc.UpdateEvent(context.Background(), updateEventDto, int32(1), int32(1))

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestListEvents(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		expectedEvents := []database.Event{
			{ID: 1, Name: "Event 1"},
			{ID: 2, Name: "Event 2"},
		}

		mockRepo.On("ListEvents", mock.Anything).Return(expectedEvents, nil)

		events, err := svc.ListEvents(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 2, len(events))
		mockRepo.AssertExpectations(t)
	})
}

func TestGetEvent(t *testing.T) {
	t.Run("Event Not Found", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("GetEvent", mock.Anything, int32(99)).
			Return(database.Event{}, sql.ErrNoRows)

		_, err := svc.GetEvent(context.Background(), 99)

		assert.ErrorIs(t, err, util.ErrEventNotFound)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("GetEvent", mock.Anything, int32(1)).
			Return(database.Event{ID: 1, Name: "Test Event"}, nil)

		event, err := svc.GetEvent(context.Background(), int32(1))

		assert.NoError(t, err)
		assert.Equal(t, int32(1), event.ID)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteEvent(t *testing.T) {
	t.Run("Unauthorized Access", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("GetEvent", mock.Anything, int32(1)).
			Return(database.Event{ID: 1, OwnerID: 20}, nil)

		err := svc.DeleteEvent(context.Background(), int32(1), int32(10))

		assert.ErrorIs(t, err, util.ErrOwnerRestrictedAction)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("GetEvent", mock.Anything, int32(1)).
			Return(database.Event{ID: 1, OwnerID: 10}, nil)

		mockRepo.On("DeleteEvent", mock.Anything, int32(1)).
			Return(nil)

		err := svc.DeleteEvent(context.Background(), int32(1), int32(10))

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestReserveTicket(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("AddAttendee", mock.Anything, mock.Anything).
			Return(new(mockDbResult), nil)

		err := svc.ReserveTicket(context.Background(), int32(1), int32(10))

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestRevokeTicket(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("RemoveAttendee", mock.Anything, mock.Anything).
			Return(nil)

		err := svc.RevokeTicket(context.Background(), int32(1), int32(10))

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetEventAttendees(t *testing.T) {
	t.Run("Unauthorized Access", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("GetEvent", mock.Anything, int32(10)).
			Return(database.Event{ID: 10, OwnerID: 1}, nil)

		_, err := svc.GetEventAttendees(context.Background(), int32(2), int32(10))

		assert.ErrorIs(t, err, util.ErrOwnerRestrictedAction)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockEventRepo)
		svc := NewEventService(mockRepo)

		mockRepo.On("GetEvent", mock.Anything, int32(10)).
			Return(database.Event{ID: 10, OwnerID: 1}, nil)

		expectedAttendees := []database.GetEventAttendeesRow{
			{Name: "User 1", Email: "user1@example.com"},
			{Name: "User 2", Email: "user2@example.com"},
		}
		mockRepo.On("GetEventAttendees", mock.Anything, int32(10)).
			Return(expectedAttendees, nil)

		attendees, err := svc.GetEventAttendees(context.Background(), 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(attendees))
		mockRepo.AssertExpectations(t)
	})
}
