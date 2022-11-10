package service

import (
	"context"

	"github.com/webdevelop-pro/go-clean-arch/internal/domain/todo"
)

type Application interface {
	GetTodos(ctx context.Context) ([]todo.Task, error)
	CreateTodo(ctx context.Context, task *todo.Task) error
	UpdateTodo(ctx context.Context, id int, task todo.Task) error
	DeleteTodo(ctx context.Context, id int) error
}
