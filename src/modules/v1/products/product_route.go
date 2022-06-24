package products

import (
	"github.com/depri11/e-commerce/src/middleware"
	"github.com/depri11/e-commerce/src/modules/v1/users"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	c := db.Collection("products")
	u := db.Collection("users")

	repository := NewRepository(c)
	userRepository := users.NewRepository(u)
	service := NewService(repository, userRepository)
	handler := NewHandler(service)

	e.GET("/products", handler.QueryProducts)
	e.GET("/products/all", handler.GetAllProducts, middleware.CheckAuth)
	e.GET("/products/:id", handler.GetProductDetails)
	e.PUT("/product/:id/review", handler.CreateReview, middleware.CheckAuth)

	e.GET("/admin/products", handler.GetAllProducts, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.POST("/admin/product/new", handler.CreateProduct, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.PUT("/admin/product/upload", handler.UploadImages, middleware.CheckAuth)
	e.PUT("/admin/product/:id", handler.UpdateProduct, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.DELETE("/admin/product/:id", handler.DeletProduct, middleware.CheckAuth, middleware.CheckRoleAdmin)

	e.GET("/admin/reviews", handler.GetAllReviewByProductId, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.DELETE("/admin/review", handler.DeleteReview, middleware.CheckAuth, middleware.CheckRoleAdmin)
}
