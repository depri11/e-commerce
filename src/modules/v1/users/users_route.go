package users

import (
	"github.com/depri11/e-commerce/src/middleware"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("users")

	repository := NewRepository(c)
	service := NewService(repository)
	handler := NewHandler(service)

	e.POST("/register", handler.Register)

	e.GET("/me", handler.GetUserDetails, middleware.CheckAuth)
	e.PUT("/me/update", handler.UpdateProfile, middleware.CheckAuth)
	e.PUT("/me/avatar", handler.UploadAvatar, middleware.CheckAuth)

	e.POST("/password/forgot", handler.ForgotPassword)
	e.PUT("/password/reset/:token", handler.ResetPassword)

	e.GET("/admin/users", handler.FindAll, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.GET("/admin/user/:id", handler.GetUserID, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.PUT("/admin/user/:id", handler.UpdateUser, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.DELETE("/admin/user/:id", handler.DeletUser, middleware.CheckAuth, middleware.CheckRoleAdmin)
}
