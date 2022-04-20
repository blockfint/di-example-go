package controller

import (
	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/blockfint/di-example-go/app/repository"
	"go.uber.org/zap"
)

type TodoController struct {
	repo   dbRepository[model.Todo]
	logger *zap.SugaredLogger
}

func NewTodoController(repo *repository.TodoRepository, logger *zap.SugaredLogger) *TodoController {
	return &TodoController{repo, logger}
}

func (con *TodoController) List() (*[]model.Todo, error) {
	return con.repo.List()
}

func (con *TodoController) Create(todo *model.Todo) error {
	return con.repo.Create(todo)
}

func (con *TodoController) FindByID(todoID uint) (*model.Todo, error) {
	return con.repo.FindByID(todoID)
}
