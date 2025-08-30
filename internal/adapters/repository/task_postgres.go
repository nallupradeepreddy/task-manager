package repository

import (
	"github.com/nallupradeepreddy/task-manager/internal/core/domain"
	"gorm.io/gorm"
)

type TaskPostgresRepository struct {
	db *gorm.DB
}

func NewTaskPostgresRepository(db *gorm.DB) *TaskPostgresRepository {
	return &TaskPostgresRepository{db: db}
}

func (r *TaskPostgresRepository) CreateTask(task *domain.Task) error {
	return r.db.Create(task).Error
}
