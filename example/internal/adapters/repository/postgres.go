package repository

import (
	"context"

	"github.com/webdevelop-pro/go-clean-arch/internal/domain/todo"
	"github.com/webdevelop-pro/go-common/configurator"
	"github.com/webdevelop-pro/go-common/db"
)

// Repository is a struct for database operations
type Repository struct {
	db *db.DB
}

// New returns a new Repository
func New(c *configurator.Configurator) *Repository {
	return &Repository{
		db: db.New(c),
	}
}

// GetTodos returns all tasks
func (r *Repository) GetTodos(ctx context.Context) ([]todo.Task, error) {
	const query = `SELECT id, title, description, status FROM tasks;`

	var tasks []todo.Task

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := todo.Task{}

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// CreateTodo creates a new task
func (r *Repository) CreateTodo(ctx context.Context, task *todo.Task) error {
	const query = `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id;`

	err := r.db.QueryRow(ctx, query, task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTodo updates a task
func (r *Repository) UpdateTodo(ctx context.Context, id int, task todo.Task) error {
	const query = `UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4;`

	_, err := r.db.Exec(ctx, query, task.Title, task.Description, task.Status, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTodo deletes a task
func (r *Repository) DeleteTodo(ctx context.Context, id int) error {
	const query = `DELETE FROM tasks WHERE id = $1;`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
