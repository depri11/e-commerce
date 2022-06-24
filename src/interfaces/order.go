package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/input"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	FindAll() ([]*models.Order, error)
	FindByID(id string) (*models.Order, error)
	FindByUserID(id string) ([]*models.Order, error)
	Insert(order *models.Order) (*mongo.InsertOneResult, error)
	Update(id string, order *models.Order) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type OrderService interface {
	GetAllOrders() (*helper.Res, error)
	FindByID(id string) (*helper.Res, error)
	FindByUserID(id string) (*helper.Res, error)
	Create(id string, input *input.CreateOrderInput) (*helper.Res, error)
	Update(id string, order *models.Order) (*helper.Res, error)
	Delele(id string) (*helper.Res, error)
}
