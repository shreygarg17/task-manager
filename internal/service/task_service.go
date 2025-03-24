package service

import (
	"context"
	"errors"
	"time"

	"github.com/michaelorina/go-tasker/internal/model"
	"github.com/michaelorina/go-tasker/internal/repository"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	return s.Repo.FetchAll(ctx)
}

func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*model.Task, error) {
	return s.Repo.FetchByID(ctx, id)
}

func (s *TaskService) CreateTask(ctx context.Context, task model.Task) (*model.Task, error) {
	if task.Title == "" {
		return nil, errors.New("title is required")
	}

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	task.Completed = false // Default state

	return s.Repo.Create(ctx, task)
}

func (s *TaskService) UpdateTask(ctx context.Context, id string, task model.Task) (*model.Task, error) {
	if task.Title == "" {
		return nil, errors.New("title is required")
	}

	task.UpdatedAt = time.Now()
	return s.Repo.Update(ctx, id, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}
