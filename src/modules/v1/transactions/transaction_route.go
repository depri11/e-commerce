package transaction

import (
	"github.com/depri11/e-commerce/src/middleware"
	payment "github.com/depri11/e-commerce/src/modules/v1/payments"
	"github.com/depri11/e-commerce/src/modules/v1/products"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("transactions")

	repository := NewRepository(c)
	repositoryProduct := products.NewRepository(db.Collection("products"))
	paymentService := payment.NewService(repository, repositoryProduct)
	service := NewService(repository, repositoryProduct, paymentService)
	handler := NewHandler(service, paymentService)

	e.GET("/product/:id/transaction", handler.GetProductTransactions, middleware.CheckAuth)

	e.GET("/transactions", handler.GetUserTransactions, middleware.CheckAuth)
	e.POST("/transaction/new", handler.CreateTransaction, middleware.CheckAuth)
	e.POST("/transaction/notification", handler.GetNotification, middleware.CheckAuth)

}
