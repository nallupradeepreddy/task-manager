package service

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/nallupradeepreddy/task-manager/internal/core/domain"
	"github.com/nallupradeepreddy/task-manager/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) RegisterUser(email, password string) (*domain.User, error) {
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}
	if existing, _ := s.repo.GetUserByEmail(email); existing != nil {
		return nil, errors.New("email already registered")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
