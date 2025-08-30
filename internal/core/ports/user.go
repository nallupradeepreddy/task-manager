package ports

import "github.com/nallupradeepreddy/task-manager/internal/core/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
}

type UserService interface {
	RegisterUser(email, password string) (*domain.User, error)
}
