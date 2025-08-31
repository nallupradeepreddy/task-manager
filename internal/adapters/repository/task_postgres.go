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

func (r *TaskPostgresRepository) FetchTasks(userID, status string) ([]*domain.Task, error) {
	var tasks []*domain.Task
	query := r.db.Where("user_id = ?", userID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskPostgresRepository) UpdateTask(task *domain.Task) (*domain.Task, error) {
	if err := r.db.Save(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}
