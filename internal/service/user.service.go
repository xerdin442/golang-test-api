package service

import (
	"context"
	"errors"

	database "github.com/xerdin442/api-practice/internal/adapters/generated"
	"github.com/xerdin442/api-practice/internal/api/dto"
	"github.com/xerdin442/api-practice/internal/api/middleware"
	repo "github.com/xerdin442/api-practice/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repo.UserRepoInterface
}

func NewUserService(r repo.UserRepoInterface) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Signup(ctx context.Context, dto dto.SignupRequest) (database.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return database.User{}, err
	}

	args := database.CreateUserParams{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: string(hashedPassword),
	}

	result, err := s.repo.CreateUser(ctx, args)
	if err != nil {
		return database.User{}, err
	}

	userID, _ := result.LastInsertId()
	return s.repo.GetUserByID(ctx, int32(userID))
}

func (s *UserService) Login(ctx context.Context, dto dto.LoginRequest) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		return "", errors.New("Invalid email address")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return "", errors.New("Invalid password")

		default:
			return "", err
		}
	}

	return middleware.GenerateToken(user.ID)
}

func (s *UserService) Logout(ctx context.Context, userID int32) error {
	return nil
}
