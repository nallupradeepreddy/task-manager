package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/nallupradeepreddy/task-manager/internal/core/domain"
	"github.com/nallupradeepreddy/task-manager/internal/core/ports"
)

type TaskServiceImpl struct {
	repo ports.TaskRepository
}

func NewTaskService(repo ports.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{repo: repo}
}

func (s *TaskServiceImpl) CreateTask(userID, title, description string) (*domain.Task, error) {
	if title == "" {
		return nil, ErrTitleRequired
	}
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	task := &domain.Task{
		ID:          uuid.New(),
		UserID:      uid,
		Title:       title,
		Description: description,
		Status:      "to_do",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.repo.CreateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

var ErrTitleRequired = &TaskError{"title is required"}

type TaskError struct {
	msg string
}

func (e *TaskError) Error() string {
	return e.msg
}
