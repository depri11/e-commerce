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

	e.GET("/products", handler.GetProducts)
	e.GET("/products/:id", handler.GetProduct)
	e.POST("/products", handler.CreateProduct)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:od", handler.DeletProduct)
}
