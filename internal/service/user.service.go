package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hibiken/asynq"
	database "github.com/xerdin442/api-practice/internal/adapters/generated"
	"github.com/xerdin442/api-practice/internal/api/dto"
	"github.com/xerdin442/api-practice/internal/api/middleware"
	"github.com/xerdin442/api-practice/internal/env"
	repo "github.com/xerdin442/api-practice/internal/repository"
	"github.com/xerdin442/api-practice/internal/tasks"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(r repo.UserRepo) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Signup(ctx context.Context, dto dto.SignupRequest, queue *asynq.Client) (database.User, error) {
	user, _ := s.repo.GetUserByEmail(ctx, dto.Email)
	if user.Email == dto.Email {
		return database.User{}, ErrEmailAlreadyExists
	}

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

	// Parse email template
	templateData := &tasks.OnboardingTemplateData{
		Name:    dto.Name,
		Company: env.GetStr("APP_NAME"),
	}
	content, _ := tasks.ParseEmailTemplate(templateData, "templates/onboarding.html")

	// Configure email payload
	payload := &tasks.EmailPayload{
		Recipient: dto.Email,
		Subject:   "Welcome Onboard!",
		Content:   content,
	}

	// Send onboarding email to new user
	task, _ := tasks.NewEmailTask(payload)
	_, err = queue.Enqueue(task)

	userID, _ := result.LastInsertId()
	return s.repo.GetUserByID(ctx, int32(userID))
}

func (s *UserService) Login(ctx context.Context, dto dto.LoginRequest) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, dto.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrInvalidEmail
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return "", ErrInvalidPassword

		default:
			return "", err
		}
	}

	return middleware.GenerateToken(user.ID)
}

func (s *UserService) GetProfile(ctx context.Context, userID int32) (database.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}
