package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
)

type TransactionRepository interface {
	FindAll() ([]*models.Transaction, error)
	FindByID(id string) (*models.Transaction, error)
	FindByProductId(id string) (*models.Transaction, error)
	FindByUserId(id string) ([]*models.Transaction, error)
	Insert(transaction *models.Transaction) (string, error)
	Update(id string, user *models.Transaction) (*models.Transaction, error)
}

type TransactionService interface {
	GetAll() (*helper.Res, error)
	GetByID(id string) (*helper.Res, error)
	GetByProductID(id string) (*helper.Res, error)
	GetByUserID(id string) (*helper.Res, error)
	Create(id string, transaction *models.Transaction) (*helper.Res, error)
}
