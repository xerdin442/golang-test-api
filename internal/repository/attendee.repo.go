package repo

import (
	"context"
	"database/sql"

	database "github.com/xerdin442/api-practice/internal/adapters/generated"
)

type AttendeeRepoInterface interface {
	AddAttendee(ctx context.Context, arg database.AddAttendeeParams) (sql.Result, error)
	RemoveAttendee(ctx context.Context, arg database.RemoveAttendeeParams) error
	GetEventAttendees(ctx context.Context, eventID int32) ([]database.GetEventAttendeesRow, error)
}

type AttendeeRepo struct {
	q *database.Queries
}

func NewAttendeeRepository(db *sql.DB) AttendeeRepoInterface {
	repo := &AttendeeRepo{q: database.New(db)}
	return repo.q
}
