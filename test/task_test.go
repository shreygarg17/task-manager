package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/michaelorina/go-tasker/internal/model"
	"github.com/michaelorina/go-tasker/internal/service"
)

// --- Mock repository (minimal version for testing) ---

type mockTaskRepo struct {
	fakeStore map[string]model.Task
	err       error
}

func (m *mockTaskRepo) FetchAll(ctx context.Context) ([]model.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	tasks := make([]model.Task, 0, len(m.fakeStore))
	for _, task := range m.fakeStore {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *mockTaskRepo) FetchByID(ctx context.Context, id string) (*model.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	task, exists := m.fakeStore[id]
	if !exists {
		return nil, errors.New("not found")
	}
	return &task, nil
}

func (m *mockTaskRepo) Create(ctx context.Context, task model.Task) (*model.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.fakeStore[task.ID] = task
	return &task, nil
}

func (m *mockTaskRepo) Update(ctx context.Context, id string, task model.Task) (*model.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	_, exists := m.fakeStore[id]
	if !exists {
		return nil, errors.New("not found")
	}
	m.fakeStore[id] = task
	return &task, nil
}

func (m *mockTaskRepo) Delete(ctx context.Context, id string) error {
	if m.err != nil {
		return m.err
	}
	delete(m.fakeStore, id)
	return nil
}

func newMockRepo() *mockTaskRepo {
	return &mockTaskRepo{
		fakeStore: make(map[string]model.Task),
	}
}

// --- Unit Tests ---

func TestCreateTask(t *testing.T) {
	repo := newMockRepo()
	svc := service.NewTaskService(repo)

	task := model.Task{
		ID:          "task123",
		Title:       "Write Go tests",
		Description: "Create unit tests for TaskService",
	}

	ctx := context.Background()
	result, err := svc.CreateTask(ctx, task)
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	if result.Title != task.Title {
		t.Errorf("Expected title %q, got %q", task.Title, result.Title)
	}
	if result.Completed {
		t.Errorf("Expected task to be not completed by default")
	}
}

func TestGetTaskByID(t *testing.T) {
	repo := newMockRepo()
	svc := service.NewTaskService(repo)

	task := model.Task{
		ID:    "abc123",
		Title: "Sample Task",
	}
	repo.fakeStore[task.ID] = task

	ctx := context.Background()
	result, err := svc.GetTaskByID(ctx, task.ID)
	if err != nil {
		t.Fatalf("GetTaskByID failed: %v", err)
	}
	if result.ID != task.ID {
		t.Errorf("Expected ID %q, got %q", task.ID, result.ID)
	}
}

func TestUpdateTask(t *testing.T) {
	repo := newMockRepo()
	svc := service.NewTaskService(repo)

	initial := model.Task{
		ID:    "xyz456",
		Title: "Old Title",
	}
	repo.fakeStore[initial.ID] = initial

	updated := model.Task{
		Title:       "Updated Title",
		Description: "Updated Desc",
	}

	ctx := context.Background()
	result, err := svc.UpdateTask(ctx, initial.ID, updated)
	if err != nil {
		t.Fatalf("UpdateTask failed: %v", err)
	}
	if result.Title != updated.Title {
		t.Errorf("Expected title %q, got %q", updated.Title, result.Title)
	}
}

func TestDeleteTask(t *testing.T) {
	repo := newMockRepo()
	svc := service.NewTaskService(repo)

	task := model.Task{
		ID:    "del-001",
		Title: "To Delete",
	}
	repo.fakeStore[task.ID] = task

	ctx := context.Background()
	err := svc.DeleteTask(ctx, task.ID)
	if err != nil {
		t.Fatalf("DeleteTask failed: %v", err)
	}

	if _, exists := repo.fakeStore[task.ID]; exists {
		t.Errorf("Expected task %q to be deleted", task.ID)
	}
}
