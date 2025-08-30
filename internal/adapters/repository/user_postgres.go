package repository

import (
	"gorm.io/gorm"

	"github.com/nallupradeepreddy/task-manager/internal/core/domain"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserPostgresRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
