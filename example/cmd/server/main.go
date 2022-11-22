package main

import (
	_ "github.com/webdevelop-pro/go-clean-arch/api" // import api
	"github.com/webdevelop-pro/go-clean-arch/internal/adapters/repository"
	"github.com/webdevelop-pro/go-clean-arch/internal/app"
	"github.com/webdevelop-pro/go-clean-arch/internal/domain/todo"
	"github.com/webdevelop-pro/go-clean-arch/internal/ports"
	"github.com/webdevelop-pro/go-clean-arch/internal/service"
	"github.com/webdevelop-pro/go-common/configurator"
	"github.com/webdevelop-pro/go-common/logger"
	"github.com/webdevelop-pro/go-common/server"
	"go.uber.org/fx"
)

// @title Todo API
// @version 0.1

// @contact.name Artem Tiumentcev
// @contact.url https://webdevelop.pro/
// @contact.email artem@webdevelop.pro

// @host localhost:8085
// @BasePath /
// @query.collection.format multi

// @schemes http https
func main() {
	fx.New(
		// Logger for fx
		fx.Provide(
			// Default logger
			logger.NewDefault,
			// Configurator helps to each component get a configuration
			configurator.New,
			// Repository
			repository.New,
			// http handler
			server.New,
			// Bind implementation to interface
			func(repo *repository.Repository) todo.Repository {
				return repo
			},
			// Init application
			app.New,
			// Bind implementation to interface
			func(app *app.App) service.Application {
				return app
			},
			// Init http handler
			ports.NewTodoHandler,
		),
		fx.Invoke(
			// Init a handler
			func(handler *ports.TodoHandler, srv *server.HttpServer) {
				for _, route := range handler.GetRoutes() {
					srv.AddRoute(route)
				}
			},
			// Start a server
			server.StartServer,
		),
	).Run()
}
