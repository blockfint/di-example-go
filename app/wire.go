//go:build wireinject
// +build wireinject

package app

import (
	"github.com/blockfint/di-example-go/app/controller"
	"github.com/blockfint/di-example-go/app/db"
	"github.com/blockfint/di-example-go/app/handler"
	"github.com/blockfint/di-example-go/app/logger"
	"github.com/blockfint/di-example-go/app/repository"
	"github.com/blockfint/di-example-go/app/server"
	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewTodoController,
	controller.NewCustomerController,
)

var repositorySet = wire.NewSet(
	repository.NewTodoRepository,
	repository.NewCustomerRepository,
)

func InitializeApplication() (*Application, error) {
	wire.Build(
		NewApplication,
		server.New,
		handler.New,
		controllerSet,
		repositorySet,
		db.New,
		logger.New,
		logger.NewZapLogger,
	)

	return nil, nil
}
