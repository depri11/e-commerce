package users

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("users")

	repository := NewRepository(c)
	service := NewService(repository)
	handler := NewHandler(service)

	e.GET("/users", handler.FindAll)
	e.POST("/users", handler.Register)
}
