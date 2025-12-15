package service

import repo "github.com/xerdin442/api-practice/internal/repository"

type UserService struct {
	repo repo.UserRepoInterface
}

func NewUserService(r repo.UserRepoInterface) *UserService {
	return &UserService{repo: r}
}
