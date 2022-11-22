package ports

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/webdevelop-pro/go-clean-arch/internal/domain/todo"
	"github.com/webdevelop-pro/go-clean-arch/internal/service"
	"github.com/webdevelop-pro/go-common/server"
)

type TodoHandler struct {
	app service.Application
}

func NewTodoHandler(app service.Application) *TodoHandler {
	return &TodoHandler{app}
}

func (h *TodoHandler) GetRoutes() []server.Route {
	return []server.Route{
		{
			Method: http.MethodGet,
			Path:   "/todos",
			Handle: h.GetTodos,
		},
		{
			Method: http.MethodPost,
			Path:   "/todos",
			Handle: h.CreateTodo,
		},
		{
			Method: http.MethodPatch,
			Path:   "/todos/:id",
			Handle: h.UpdateTodo,
		},
		{
			Method: http.MethodDelete,
			Path:   "/todos/:id",
			Handle: h.DeleteTodo,
		},
		{
			Method: http.MethodGet,
			Path:   "/docs/*",
			Handle: echoSwagger.WrapHandler,
		},
	}
}

// @Summary Get all tasks
// @Accept json
// @Produce json
// @Success 200 {object} todo.Tasks
// @Router /todos [get]
func (h *TodoHandler) GetTodos(c echo.Context) error {
	todos, err := h.app.GetTodos(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todos)
}

// @Summary Create a task
// @Accept json
// @Produce json
// @Param body body todo.Task true " "
// @Success 200 {object} todo.Tasks
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var reqData todo.Task

	err := c.Bind(&reqData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.app.CreateTodo(c.Request().Context(), &reqData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, reqData)
}

// @Summary Update a task
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param body body todo.Task true " "
// @Success 200 {object} todo.Tasks
// @Router /todos/{id} [patch]
func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	var reqData todo.Task

	id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = c.Bind(&reqData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.app.UpdateTodo(c.Request().Context(), intID, reqData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, reqData)
}

// @Summary Delete a task
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 ""
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.app.DeleteTodo(c.Request().Context(), intID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
