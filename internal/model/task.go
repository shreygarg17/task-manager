package model

import "time"

type Task struct {
	ID          string    `json:"id"`           // UUID or DB-generated
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
