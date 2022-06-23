package interfaces

import (
	"mime/multipart"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/input"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*input.RespUser, error)
	Insert(user *models.User) (*mongo.InsertOneResult, error)
	Update(id string, user *models.User) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
	FindByResetPassToken(token string) (*models.User, error)
}

type UserService interface {
	FindAll() (*helper.Res, error)
	GetUserID(id string) (*helper.Res, error)
	Insert(input *input.UserInput) (*helper.Res, error)
	Update(id string, user *input.UserInput) (*helper.Res, error)
	Delete(id string) (*helper.Res, error)
	ForgotPassword(input *input.ForgotPasswordInput) (*helper.Res, error)
	ResetPassword(token string, user *models.User) (*helper.Res, error)
	UploadAvatar(id string, file multipart.File, handle *multipart.FileHeader) (*helper.Res, error)
}
