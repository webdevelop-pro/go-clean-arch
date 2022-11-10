package todo

import (
	"context"
)

// Repository is an interface for database operations
type Repository interface {
	GetTodos(ctx context.Context) ([]Task, error)
	CreateTodo(ctx context.Context, task *Task) error
	UpdateTodo(ctx context.Context, id int, task Task) error
	DeleteTodo(ctx context.Context, id int) error
}
