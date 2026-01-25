package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	database "github.com/xerdin442/api-practice/internal/adapters/generated"
	"github.com/xerdin442/api-practice/internal/api/dto"
	"github.com/xerdin442/api-practice/internal/config"
	"github.com/xerdin442/api-practice/internal/util"
	"golang.org/x/crypto/bcrypt"
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

type mockTasksClient struct {
	mock.Mock
}

func (m *mockTasksClient) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	args := m.Called(task, opts)
	return args.Get(0).(*asynq.TaskInfo), args.Error(1)
}

func TestSignup(t *testing.T) {
	testCfg := &config.Config{
		DefaultProfileImage: "default_profile_image",
	}

	mockClient := new(mockTasksClient)

	signupDto := dto.SignupRequest{
		Email:    "test@example.com",
		Password: "Password123",
		Name:     "Test User",
	}

	t.Run("Email Already Exists", func(t *testing.T) {
		mockRepo := new(mockUserRepo)
		svc := NewUserService(mockRepo, testCfg)

		mockRepo.On("GetUserByEmail", mock.Anything, signupDto.Email).
			Return(database.User{Email: signupDto.Email}, nil)

		_, err := svc.Signup(context.Background(), signupDto, mockClient)

		assert.ErrorIs(t, err, util.ErrEmailAlreadyExists)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockUserRepo)
		svc := NewUserService(mockRepo, testCfg)

		mockRepo.On("GetUserByEmail", mock.Anything, signupDto.Email).
			Return(database.User{}, nil)

		mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(p database.CreateUserParams) bool {
			return p.Email == signupDto.Email &&
				p.Name == signupDto.Name &&
				p.ProfileImage.String == testCfg.DefaultProfileImage
		})).Return(new(mockResult), nil)

		mockRepo.On("GetUserByID", mock.Anything, mock.Anything).
			Return(database.User{Email: signupDto.Email}, nil)

		taskInfo := &asynq.TaskInfo{
			ID:      "test-id",
			Queue:   "default",
			Payload: []byte("test-payload"),
		}
		mockClient.On("Enqueue", mock.Anything, mock.Anything).Return(taskInfo, nil)

		_, err := svc.Signup(context.Background(), signupDto, mockClient)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	testCfg := &config.Config{
		JwtSecret: "jwt_secret_key",
	}

	loginDto := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "Password123",
	}

	t.Run("Invalid Email", func(t *testing.T) {
		mockRepo := new(mockUserRepo)
		svc := NewUserService(mockRepo, testCfg)

		mockRepo.On("GetUserByEmail", mock.Anything, loginDto.Email).
			Return(database.User{Email: "xerdin442@example.com"}, sql.ErrNoRows)

		_, err := svc.Login(context.Background(), loginDto)

		assert.ErrorIs(t, err, util.ErrInvalidEmail)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		mockRepo := new(mockUserRepo)
		svc := NewUserService(mockRepo, testCfg)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Password456"), bcrypt.DefaultCost)

		mockRepo.On("GetUserByEmail", mock.Anything, loginDto.Email).
			Return(database.User{Email: loginDto.Email, Password: string(hashedPassword)}, nil)

		_, err := svc.Login(context.Background(), loginDto)

		assert.ErrorIs(t, err, util.ErrInvalidPassword)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mockUserRepo)
		svc := NewUserService(mockRepo, testCfg)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(loginDto.Password), bcrypt.DefaultCost)

		mockRepo.On("GetUserByEmail", mock.Anything, loginDto.Email).
			Return(database.User{Email: loginDto.Email, Password: string(hashedPassword)}, nil)

		_, err := svc.Login(context.Background(), loginDto)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
