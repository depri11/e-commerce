package products

import (
	"github.com/depri11/e-commerce/src/middleware"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("products")

	repository := NewRepository(c)
	service := NewService(repository)
	handler := NewHandler(service)

	e.GET("/products", handler.QueryProducts)
	e.GET("/products/all", handler.GetAllProducts, middleware.CheckAuth)

	e.GET("/products/:id", handler.GetProductDetails)
	e.POST("/products", handler.CreateProduct, middleware.CheckAuth)
	e.PUT("/products/:id", handler.UpdateProduct, middleware.CheckAuth)
	e.DELETE("/products/:id", handler.DeletProduct, middleware.CheckAuth)
}
