package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/michaelorina/go-tasker/internal/config"
	"github.com/michaelorina/go-tasker/internal/model"
	"github.com/michaelorina/go-tasker/internal/repository/postgres"
	"github.com/michaelorina/go-tasker/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: taskctl [list|create|delete|help]")
		return
	}

	cfg := config.Load()

	db, err := config.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	repo := postgres.NewTaskRepo(db)
	svc := service.NewTaskService(repo)

	cmd := os.Args[1]

	switch cmd {
	case "list":
		tasks, err := svc.GetAllTasks(context.Background())
		if err != nil {
			log.Fatal("Fetch failed:", err)
		}
		for _, task := range tasks {
			fmt.Printf("- [%s] %s (Done: %v)\n", task.ID, task.Title, task.Completed)
		}

	case "create":
		if len(os.Args) < 4 {
			fmt.Println("Usage: taskctl create <title> <description>")
			return
		}
		title := os.Args[2]
		desc := os.Args[3]

		newTask := model.Task{
			Title:       title,
			Description: desc,
		}
		task, err := svc.CreateTask(context.Background(), newTask)
		if err != nil {
			log.Fatal("Create failed:", err)
		}
		fmt.Println("✅ Task created with ID:", task.ID)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskctl delete <task_id>")
			return
		}
		id := os.Args[2]
		err := svc.DeleteTask(context.Background(), id)
		if err != nil {
			log.Fatal("Delete failed:", err)
		}
		fmt.Println("🗑️ Task deleted:", id)

	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Available commands: list | create | delete | help ")
	}
}
