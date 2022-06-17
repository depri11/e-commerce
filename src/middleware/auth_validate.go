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

		_, err := c.Cookie("token")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		checkToken, err := helper.CheckToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		c.Request().Header.Set("user_id", checkToken.Id)
		c.Request().Header.Set("user_name", checkToken.Name)
		c.Request().Header.Set("user_email", checkToken.Email)
		c.Request().Header.Set("user_role", checkToken.Role)

		return next(c)
	}
}
