package repository

import (
	"errors"

	"github.com/blockfint/di-example-go/app/db/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TodoRepository struct {
	gormDB *gorm.DB
	logger *zap.SugaredLogger
}

func NewTodoRepository(gormDB *gorm.DB, logger *zap.SugaredLogger) *TodoRepository {
	return &TodoRepository{gormDB, logger}
}

func (repo *TodoRepository) List() (*[]model.Todo, error) {
	var todos []model.Todo
	if result := repo.gormDB.Find(&todos); result.Error != nil {
		return nil, result.Error
	}

	return &todos, nil
}

func (repo *TodoRepository) Create(todo *model.Todo) error {
	result := repo.gormDB.Create(&todo)

	return result.Error
}

func (repo *TodoRepository) FindByID(todoID uint) (*model.Todo, error) {
	var todo model.Todo
	if result := repo.gormDB.Where(todoID).First(&todo); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &todo, nil
}
