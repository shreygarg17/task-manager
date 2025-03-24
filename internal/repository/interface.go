package repository

import (
	"context"

	"github.com/michaelorina/go-tasker/internal/model"
)

type TaskRepository interface {
	FetchAll(ctx context.Context) ([]model.Task, error)
	FetchByID(ctx context.Context, id string) (*model.Task, error)
	Create(ctx context.Context, task model.Task) (*model.Task, error)
	Update(ctx context.Context, id string, task model.Task) (*model.Task, error)
	Delete(ctx context.Context, id string) error
}
