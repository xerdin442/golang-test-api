package service

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	database "github.com/xerdin442/api-practice/internal/adapters/generated"
	"github.com/xerdin442/api-practice/internal/api/dto"
)

type mockResult struct{}

func (m *mockResult) LastInsertId() (int64, error) { return 1, nil }
func (m *mockResult) RowsAffected() (int64, error) { return 1, nil }

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

func TestSignup(t *testing.T) {
	os.Setenv("JWT_SECRET", "jwt_secret_key")
	os.Setenv("APP_NAME", "Test App")
	os.Setenv("DEFAULT_PROFILE_IMAGE", "default_profile_image")

	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		signupDto := dto.SignupRequest{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "Test User",
		}

		mockRepo.On("GetUserByEmail", mock.Anything, signupDto.Email).
			Return(database.User{}, sql.ErrNoRows)

		mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(p database.CreateUserParams) bool {
			return p.Email == signupDto.Email &&
				p.Name == signupDto.Name &&
				p.ProfileImage.String == secrets.DefaultProfileImage
		})).Return(new(mockResult), nil)

		mockRepo.On("GetUserByID", mock.Anything, mock.Anything).
			Return(database.User{Email: signupDto.Email}, nil)

		var taskClient *asynq.Client
		_, err := svc.Signup(context.Background(), signupDto, taskClient)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
