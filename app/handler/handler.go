package handler

import (
	"github.com/blockfint/di-example-go/app/controller"
	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type handlerController[M interface{}] interface {
	List() (*[]M, error)
	FindByID(id uint) (*M, error)
	Create(m *M) error
}

type Handler struct {
	todoController     handlerController[model.Todo]
	customerController handlerController[model.Customer]
	logger             *zap.SugaredLogger
}

func New(
	todoController *controller.TodoController,
	customerController *controller.CustomerController,
	logger *zap.SugaredLogger,
) *Handler {
	return &Handler{
		todoController,
		customerController,
		logger,
	}
}

type route struct {
	Path        string
	Method      string
	HandlerFunc func(echo.Context) error
}

type group struct {
	Prefix      string
	middlewares []echo.MiddlewareFunc
	routes      []route
}
