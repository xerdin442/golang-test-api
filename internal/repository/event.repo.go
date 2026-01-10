package repo

import (
	"context"
	"database/sql"

	database "github.com/xerdin442/api-practice/internal/adapters/generated"
)

type EventRepoInterface interface {
	CreateEvent(ctx context.Context, arg database.CreateEventParams) (sql.Result, error)
	GetEvent(ctx context.Context, id int32) (database.Event, error)
	UpdateEvent(ctx context.Context, arg database.UpdateEventParams) (sql.Result, error)
	DeleteEvent(ctx context.Context, id int32) error
	ListEvents(ctx context.Context) ([]database.Event, error)
	AddAttendee(ctx context.Context, arg database.AddAttendeeParams) (sql.Result, error)
	RemoveAttendee(ctx context.Context, arg database.RemoveAttendeeParams) error
	GetEventAttendees(ctx context.Context, eventID int32) ([]database.GetEventAttendeesRow, error)
}

type EventRepo struct {
	q *database.Queries
}

func NewEventRepository(db *sql.DB) EventRepoInterface {
	repo := EventRepo{q: database.New(db)}
	return repo.q
}
