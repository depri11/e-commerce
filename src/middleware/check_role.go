package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckRoleAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Request().Header.Get("user_role")
		if role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "Access danied")
		}

		return next(c)
	}
}
