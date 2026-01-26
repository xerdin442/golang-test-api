package service

import (
	"context"
	"database/sql"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/mock"
	database "github.com/xerdin442/api-practice/internal/adapters/generated"
)

type mockDbResult struct{}

func (m *mockDbResult) LastInsertId() (int64, error) { return 1, nil }
func (m *mockDbResult) RowsAffected() (int64, error) { return 1, nil }

type mockTasksClient struct {
	mock.Mock
}

func (m *mockTasksClient) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	args := m.Called(task, opts)
	return args.Get(0).(*asynq.TaskInfo), args.Error(1)
}

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) CreateUser(ctx context.Context, params database.CreateUserParams) (sql.Result, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(sql.Result), args.Error(1)
}

func (m *mockUserRepo) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *mockUserRepo) GetUserByID(ctx context.Context, id int32) (database.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(database.User), args.Error(1)
}

type mockEventRepo struct {
	mock.Mock
}

func (m *mockEventRepo) CreateEvent(ctx context.Context, params database.CreateEventParams) (sql.Result, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(sql.Result), args.Error(1)
}

func (m *mockEventRepo) GetEvent(ctx context.Context, id int32) (database.Event, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(database.Event), args.Error(1)
}

func (m *mockEventRepo) UpdateEvent(ctx context.Context, params database.UpdateEventParams) (sql.Result, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(sql.Result), args.Error(1)
}

func (m *mockEventRepo) DeleteEvent(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockEventRepo) ListEvents(ctx context.Context) ([]database.Event, error) {
	args := m.Called(ctx)
	return args.Get(0).([]database.Event), args.Error(1)
}

func (m *mockEventRepo) AddAttendee(ctx context.Context, params database.AddAttendeeParams) (sql.Result, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(sql.Result), args.Error(1)
}

func (m *mockEventRepo) RemoveAttendee(ctx context.Context, params database.RemoveAttendeeParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *mockEventRepo) GetEventAttendees(ctx context.Context, eventID int32) ([]database.GetEventAttendeesRow, error) {
	args := m.Called(ctx, eventID)
	return args.Get(0).([]database.GetEventAttendeesRow), args.Error(1)
}
