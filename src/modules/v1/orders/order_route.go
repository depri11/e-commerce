package orders

import (
	"github.com/depri11/e-commerce/src/middleware"
	"github.com/depri11/e-commerce/src/modules/v1/payments"
	"github.com/depri11/e-commerce/src/modules/v1/products"
	"github.com/depri11/e-commerce/src/modules/v1/users"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(e *echo.Group, db *mongo.Database) {
	o := db.Collection("orders")
	p := db.Collection("products")
	u := db.Collection("users")

	repository := NewRepository(o)
	productRepository := products.NewRepository(p)
	userRepository := users.NewRepository(u)
	paymentService := payments.NewService(repository, productRepository)
	service := NewService(repository, userRepository, paymentService)
	handler := NewHandler(service)

	e.POST("/order/new", handler.NewOrder, middleware.CheckAuth)
	e.GET("/order/:id", handler.GetOrderDetails)
	e.GET("/orders/me", handler.MyOrders, middleware.CheckAuth)

	e.DELETE("/admin/order/:id", handler.DeleteOrder, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.GET("/admin/orders", handler.GetAllOrders, middleware.CheckAuth, middleware.CheckRoleAdmin)
	e.POST("/admin/order/:id", handler.UpdateOrder, middleware.CheckAuth, middleware.CheckRoleAdmin)

}
