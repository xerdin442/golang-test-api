package service

import repo "github.com/xerdin442/api-practice/internal/repository"

type AttendeeService struct {
	repo repo.AttendeeRepoInterface
}

func NewAttendeeService(r repo.AttendeeRepoInterface) *AttendeeService {
	return &AttendeeService{repo: r}
}
