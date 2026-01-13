package repo

import (
	"context"
	"database/sql"

	database "github.com/xerdin442/api-practice/internal/adapters/generated"
)

type UserRepo interface {
	CreateUser(ctx context.Context, arg database.CreateUserParams) (sql.Result, error)
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
	GetUserByID(ctx context.Context, id int32) (database.User, error)
}

func NewUserRepository(db *sql.DB) UserRepo {
	return database.New(db)
}
