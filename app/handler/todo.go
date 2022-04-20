package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/blockfint/di-example-go/app/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) NewTodosGroup() *group {
	return &group{
		Prefix:      "/todos",
		middlewares: []echo.MiddlewareFunc{},
		routes: []route{
			{
				Path:        "",
				Method:      http.MethodGet,
				HandlerFunc: h.ListTodos,
			},
			{
				Path:        "/:todoID",
				Method:      http.MethodGet,
				HandlerFunc: h.FindTodoByID,
			},
			{
				Path:        "",
				Method:      http.MethodPost,
				HandlerFunc: h.CreateTodo,
			},
		},
	}
}

func (h *Handler) ListTodos(c echo.Context) error {
	todos, err := h.todoController.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *Handler) FindTodoByID(c echo.Context) error {
	id64, err := strconv.ParseUint(c.Param("todoID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	todoID := uint(id64)
	todo, err := h.todoController.FindByID(todoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if todo == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *Handler) CreateTodo(c echo.Context) error {
	var todo model.Todo

	req := &createTodoRequest{}

	if err := req.bind(c, &todo); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.todoController.Create(&todo); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	utils.SetLocationHeader(c, fmt.Sprintf("/todos/%d", todo.ID))
	return c.JSON(http.StatusCreated, todo)
}
