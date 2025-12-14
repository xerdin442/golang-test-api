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
}

type EventRepo struct {
	q *database.Queries
}

func NewEventRepository(db *sql.DB) EventRepoInterface {
	repo := EventRepo{q: database.New(db)}
	return repo.q
}
