package main

import (
	"log"
	"net/http"
	"os"

	"github.com/michaelorina/go-tasker/internal/config"
	httphandler "github.com/michaelorina/go-tasker/internal/handler/http"
	"github.com/michaelorina/go-tasker/internal/repository/postgres"
	"github.com/michaelorina/go-tasker/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := config.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}
	defer db.Close()

	taskRepo := postgres.NewTaskRepo(db)
	taskService := service.NewTaskService(taskRepo)

	router := httphandler.SetupRouter(taskService)

	log.Printf("🚀 Server is running on port %s...\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
		os.Exit(1)
	}
}
