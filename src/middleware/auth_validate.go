package middleware

import (
	"net/http"
	"strings"

	"github.com/depri11/e-commerce/src/helper"
	"github.com/labstack/echo/v4"
)

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")

		token := c.Request().Header.Get("Authorization")

		if !strings.Contains(token, "Bearer") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		token = strings.TrimPrefix(token, "Bearer ")
		// token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := helper.CheckToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		c.Response().Header().Set("name", checkToken.Name)

		return next(c)
	}
}
