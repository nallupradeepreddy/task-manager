package ports

import "github.com/nallupradeepreddy/task-manager/internal/core/domain"

type TaskRepository interface {
	CreateTask(task *domain.Task) error
}

type TaskService interface {
	CreateTask(userID string, title, description string) (*domain.Task, error)
}
