package auth

import (
	"github.com/depri11/e-commerce/src/modules/v1/users"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("users")

	repository := users.NewRepository(c)
	service := NewService(repository)
	handler := NewHandler(service)

	e.POST("/login", handler.SigIn)

}
