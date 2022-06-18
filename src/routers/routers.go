package routers

import (
	"github.com/depri11/e-commerce/src/database"
	"github.com/depri11/e-commerce/src/modules/v1/auth"
	"github.com/depri11/e-commerce/src/modules/v1/products"
	transaction "github.com/depri11/e-commerce/src/modules/v1/transactions"
	"github.com/depri11/e-commerce/src/modules/v1/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouters() (*echo.Echo, error) {
	e := echo.New()
	db, err := database.SetupDB()
	if err != nil {
		return nil, err
	}

	e.Use(middleware.CORS())

	api := e.Group("/api/v1")
	auth.NewRouter(api, db)
	users.NewRouter(api, db)
	products.NewRouter(api, db)
	transaction.NewRouter(api, db)

	return e, nil
}
