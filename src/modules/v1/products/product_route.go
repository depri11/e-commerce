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
	e.POST("/product/:id/review", handler.CreateReview, middleware.CheckAuth)

	e.GET("/admin/products", handler.GetAllProducts, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.POST("/admin/products/new", handler.CreateProduct, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.PUT("/admin/products/:id", handler.UpdateProduct, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.DELETE("/admin/products/:id", handler.DeletProduct, middleware.CheckAuth, middleware.CheckRoleAdmin)

	e.GET("/admin/reviews", handler.GetAllReviews, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.DELETE("/admin/reviews", handler.QueryDeleteReview, middleware.CheckAuth, middleware.CheckRoleAdmin)
}
