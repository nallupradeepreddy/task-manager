package ports

import "github.com/nallupradeepreddy/task-manager/internal/core/domain"

type TaskRepository interface {
	CreateTask(task *domain.Task) error
	FetchTasks(userID, status string) ([]*domain.Task, error)
	UpdateTask(task *domain.Task) (*domain.Task, error)
}

type TaskService interface {
	CreateTask(userID string, title, description string) (*domain.Task, error)
	FetchTasks(userID, status string) ([]*domain.Task, error)
	UpdateTask(taskID, userID, title, description, status string) (*domain.Task, error)
}
