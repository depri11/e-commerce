package routers

import (
	"github.com/depri11/e-commerce/src/database"
	"github.com/depri11/e-commerce/src/modules/v1/products"
	"github.com/depri11/e-commerce/src/modules/v1/users"
	"github.com/labstack/echo/v4"
)

func SetupRouters() (*echo.Echo, error) {
	e := echo.New()
	db, err := database.SetupDB()
	if err != nil {
		return nil, err
	}

	api := e.Group("/api/v1")
	users.NewRouter(api, db)
	products.NewRouter(api, db)

	return e, nil
}
