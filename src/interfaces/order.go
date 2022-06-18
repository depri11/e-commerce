package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	FindAll() ([]*models.Order, error)
	// FindOrderById(id string) (*models.Order, error)
	Insert(order models.Order) (*models.Order, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type OrderService interface {
	GetAllOrders() (*helper.Res, error)
	Create(order models.Order) (*helper.Res, error)
	Delele(id string) (*helper.Res, error)
}
