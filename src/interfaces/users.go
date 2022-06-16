package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
	Insert(user *models.User) (*mongo.InsertOneResult, error)
}

type UserService interface {
	FindAll() (*helper.Res, error)
	Insert(user *models.User) (*helper.Res, error)
}
