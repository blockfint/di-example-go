package repository

import (
	"errors"

	"github.com/blockfint/di-example-go/app/db/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	gormDB *gorm.DB
	logger *zap.SugaredLogger
}

func NewCustomerRepository(gormDB *gorm.DB, logger *zap.SugaredLogger) *CustomerRepository {
	return &CustomerRepository{gormDB, logger}
}

func (repo *CustomerRepository) List() (*[]model.Customer, error) {
	var customers []model.Customer
	result := repo.gormDB.Preload("IDCard.IDCardAddress").Joins("IDAddress").Joins("ContactAddress").Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customers, nil
}

func (repo *CustomerRepository) FindByID(customerID uint) (*model.Customer, error) {
	var customer model.Customer
	if result := repo.gormDB.Preload("IDCard.IDCardAddress").Joins("IDAddress").Joins("ContactAddress").Where(customerID).First(&customer); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &customer, nil
}

func (repo *CustomerRepository) Create(customer *model.Customer) error {
	result := repo.gormDB.Create(&customer)

	if saveResult := repo.gormDB.Save(&customer); saveResult.Error != nil {
		return saveResult.Error
	}

	return result.Error
}
