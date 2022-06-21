package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id string) (*models.Product, error)
	Insert(user *models.Product) (*mongo.InsertOneResult, error)
	Update(id string, product *models.Product) (*mongo.UpdateResult, error)
	UpdateMany(id []string, product *models.Product) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
	Search(page, search, sort string) ([]models.Product, error)
	// FindAllReview() ([]models.Review, error)
	// InsertReview(review *models.Review) (*mongo.InsertOneResult, error)
	DeleteReview(id string) (*mongo.DeleteResult, error)
}

type ProductService interface {
	FindAll() (*helper.Res, error)
	GetUserID(id string) (*helper.Res, error)
	Insert(product *models.Product) (*helper.Res, error)
	Update(id string, product *models.Product) (*helper.Res, error)
	Delete(id string) (*helper.Res, error)
	Search(page, search, sort string) (*helper.Res, error)
	GetReviews(id string) (*helper.Res, error)
	InsertReview(review *models.Review) (*helper.Res, error)
	DeleteReview(id string) (*helper.Res, error)
}
