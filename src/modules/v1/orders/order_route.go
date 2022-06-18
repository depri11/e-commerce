package orders

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("orders")

	repository := NewRepository(c)
	service := NewService(repository)
	handler := NewHandler(service)

	e.GET("/orders", handler.GetAllOrders)
	e.POST("/order/new", handler.Create)
	e.DELETE("/order/:id", handler.Delete)

}
