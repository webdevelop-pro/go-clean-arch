package adapters

import "github.com/webdevelop-pro/go-clean-arch/internal/domain/todo"

type Repository interface {
	todo.Repository
}
