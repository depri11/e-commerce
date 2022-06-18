package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
	Insert(user *models.User) (*models.User, error)
	Update(id string, user *models.User) (*models.User, error)
	UpdateProfile(id string, user *models.User) (*models.User, error)
	Delete(id string) (*mongo.DeleteResult, error)
	FindByResetPassToken(token string) (*models.User, error)
}

type UserService interface {
	FindAll() (*helper.Res, error)
	GetUserID(id string) (*helper.Res, error)
	Insert(user *models.User) (*helper.Res, error)
	Update(id string, user *models.User) (*helper.Res, error)
	UpdateProfile(id string, user *models.User) (*helper.Res, error)
	Delete(id string) (*helper.Res, error)
	ForgotPassword(user *models.User) (*helper.Res, error)
	ResetPassword(token string, user *models.User) (*helper.Res, error)
}
