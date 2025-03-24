package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/michaelorina/go-tasker/internal/model"
)

type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{DB: db}
}

func (r *TaskRepo) FetchAll(ctx context.Context) ([]model.Task, error) {
	rows, err := r.DB.QueryContext(ctx, `SELECT id, title, description, completed, created_at, updated_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepo) FetchByID(ctx context.Context, id string) (*model.Task, error) {
	var t model.Task
	err := r.DB.QueryRowContext(ctx, `SELECT id, title, description, completed, created_at, updated_at FROM tasks WHERE id=$1`, id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TaskRepo) Create(ctx context.Context, task model.Task) (*model.Task, error) {
	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := r.DB.ExecContext(ctx,
		`INSERT INTO tasks (id, title, description, completed, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		task.ID, task.Title, task.Description, task.Completed, task.CreatedAt, task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepo) Update(ctx context.Context, id string, task model.Task) (*model.Task, error) {
	task.UpdatedAt = time.Now()
	_, err := r.DB.ExecContext(ctx,
		`UPDATE tasks SET title=$1, description=$2, completed=$3, updated_at=$4 WHERE id=$5`,
		task.Title, task.Description, task.Completed, task.UpdatedAt, id,
	)
	if err != nil {
		return nil, err
	}
	task.ID = id
	return &task, nil
}

func (r *TaskRepo) Delete(ctx context.Context, id string) error {
	_, err := r.DB.ExecContext(ctx, `DELETE FROM tasks WHERE id=$1`, id)
	return err
}
