package repo

import (
	"context"
	"database/sql"

	database "github.com/xerdin442/api-practice/internal/adapters/generated"
)

type UserRepoInterface interface {
	CreateUser(ctx context.Context, arg database.CreateUserParams) (sql.Result, error)
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
	GetUserByID(ctx context.Context, id int32) (database.User, error)
}

type UserRepo struct {
	q *database.Queries
}

func NewUserRepository(db *sql.DB) UserRepoInterface {
	repo := &UserRepo{q: database.New(db)}
	return repo.q
}
