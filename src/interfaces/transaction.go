package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository interface {
	FindAll() ([]*models.Transaction, error)
	FindByID(id string) (*models.Transaction, error)
	FindByProductId(id string) (*models.Transaction, error)
	FindByUserId(id string) (*models.Transaction, error)
	Insert(transaction *models.Transaction) (*mongo.InsertOneResult, error)
	Update(id string, user *models.Transaction) (*models.Transaction, error)
}

type TransactionService interface {
	GetAll() (*helper.Res, error)
	GetByID(id string) (*helper.Res, error)
	Create(transaction *models.Transaction) (*helper.Res, error)
	ProcessPayment(input *models.TransactionNotification) (*helper.Res, error)
}
