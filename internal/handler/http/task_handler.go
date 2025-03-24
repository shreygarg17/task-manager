package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/michaelorina/go-tasker/internal/service"
)

// SetupRouter returns the main router with all task routes registered
func SetupRouter(taskService *service.TaskService) http.Handler {
	r := chi.NewRouter()

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", GetAllTasks(taskService))
		r.Post("/", CreateTask(taskService))
		r.Get("/{id}", GetTaskByID(taskService))
		r.Put("/{id}", UpdateTask(taskService))
		r.Delete("/{id}", DeleteTask(taskService))
	})

	return r
}

func GetAllTasks(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Replace with actual logic
		json.NewEncoder(w).Encode(map[string]string{"message": "Get all tasks"})
	}
}

func CreateTask(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Replace with actual logic
		json.NewEncoder(w).Encode(map[string]string{"message": "Create task"})
	}
}

func GetTaskByID(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		// TODO: Replace with actual logic
		json.NewEncoder(w).Encode(map[string]string{"message": "Get task by ID", "id": id})
	}
}

func UpdateTask(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		// TODO: Replace with actual logic
		json.NewEncoder(w).Encode(map[string]string{"message": "Update task", "id": id})
	}
}

func DeleteTask(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		// TODO: Replace with actual logic
		json.NewEncoder(w).Encode(map[string]string{"message": "Delete task", "id": id})
	}
}
