package controller

import (
	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/blockfint/di-example-go/app/repository"
	"go.uber.org/zap"
)

type CustomerController struct {
	repo   dbRepository[model.Customer]
	logger *zap.SugaredLogger
}

func NewCustomerController(repo *repository.CustomerRepository, logger *zap.SugaredLogger) *CustomerController {
	return &CustomerController{repo, logger}
}

func (con *CustomerController) List() (*[]model.Customer, error) {
	return con.repo.List()
}

func (con *CustomerController) FindByID(customerID uint) (*model.Customer, error) {
	return con.repo.FindByID(customerID)
}

func (con *CustomerController) Create(customer *model.Customer) error {
	return con.repo.Create(customer)
}
