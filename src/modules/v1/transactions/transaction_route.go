package transaction

import (
	"github.com/depri11/e-commerce/src/modules/v1/products"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("transactions")

	repository := NewRepository(c)
	repositoryProduct := products.NewRepository(db.Collection("products"))
	service := NewService(repository, repositoryProduct)
	handler := NewHandler(service)

	e.GET("/transactions", handler.GetTransactions)
	e.POST("/transaction/new", handler.CreateTransaction)
	e.POST("/transaction/notification", handler.GetNotification)

}
