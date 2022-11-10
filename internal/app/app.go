package app

import (
	"context"

	"github.com/webdevelop-pro/go-clean-arch/internal/domain/todo"
)

// App is an application layer
type App struct {
	repo todo.Repository
}

// New returns a new App
func New(repo todo.Repository) *App {
	return &App{repo}
}

// GetTodos returns all tasks
func (a *App) GetTodos(ctx context.Context) ([]todo.Task, error) {
	return a.repo.GetTodos(ctx)
}

// CreateTodo creates a new task
func (a *App) CreateTodo(ctx context.Context, task *todo.Task) error {
	return a.repo.CreateTodo(ctx, task)
}

// UpdateTodo updates a task
func (a *App) UpdateTodo(ctx context.Context, id int, task todo.Task) error {
	return a.repo.UpdateTodo(ctx, id, task)
}

// DeleteTodo deletes a task
func (a *App) DeleteTodo(ctx context.Context, id int) error {
	return a.repo.DeleteTodo(ctx, id)
}
