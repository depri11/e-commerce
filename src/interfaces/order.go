package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	FindAll() ([]*models.Order, error)
	FindByID(id string) (*models.Order, error)
	FindByUserID(id string) ([]*models.Order, error)
	Insert(order models.Order) (*models.Order, error)
	Update(id string, order models.Order) (*models.Order, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type OrderService interface {
	GetAllOrders() (*helper.Res, error)
	FindByID(id string) (*helper.Res, error)
	FindByUserID(id string) (*helper.Res, error)
	Create(id string, order models.Order) (*helper.Res, error)
	Update(id string, order models.Order) (*helper.Res, error)
	Delele(id string) (*helper.Res, error)
}
