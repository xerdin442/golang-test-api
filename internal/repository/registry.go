package repo

import "database/sql"

type Registry struct {
	Event EventRepoInterface
	User  UserRepoInterface
}

func NewRegistry(db *sql.DB) *Registry {
	return &Registry{
		Event: NewEventRepository(db),
		User:  NewUserRepository(db),
	}
}
