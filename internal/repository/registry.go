package repo

import "database/sql"

type Registry struct {
	Event    EventRepoInterface
	Attendee AttendeeRepoInterface
	User     UserRepoInterface
}

func NewRegistry(db *sql.DB) *Registry {
	return &Registry{
		Event:    NewEventRepository(db),
		Attendee: NewAttendeeRepository(db),
		User:     NewUserRepository(db),
	}
}
