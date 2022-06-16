package products

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("products")

	repository := NewRepository(c)
	service := NewService(repository)
	handler := NewHandler(service)

	e.GET("/products", handler.QueryProducts)
	e.GET("/products/all", handler.GetAllProducts)

	e.GET("/products/:id", handler.GetProductDetails)
	e.POST("/products", handler.CreateProduct)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:id", handler.DeletProduct)
}
