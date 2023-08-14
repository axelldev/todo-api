package model

import "time"

// Todo is a struct for todo
type Todo struct {
	ID          int       `json:"id"`          // ID is a unique identifier for a todo
	Title       string    `json:"title"`       // Title is a title of a todo
	Description string    `json:"description"` // Description is a description of a todo
	Completed   bool      `json:"completed"`   // Completed is a flag to check if a todo is completed
	CreatedAt   time.Time `json:"created_at"`  // CreatedAt is a timestamp when a todo is created
}
