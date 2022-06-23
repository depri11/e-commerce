package auth

import (
	"net/http"
	"time"

	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/input"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service interfaces.AuthService
}

func NewHandler(service interfaces.AuthService) *handler {
	return &handler{service}
}

func (h *handler) SigIn(c echo.Context) error {
	var input input.AuthLogin
	if err := c.Bind(&input); err != nil {
		return err
	}

	if err := helper.ValidationError(input); err != nil {
		res := helper.ResponseJSON("Failed to validate", 500, "error", err.Error())
		return c.JSON(500, res)
	}

	tokens, err := h.service.Login(input)
	if err != nil {
		res := helper.ResponseJSON("Faled login!", 401, "error", err.Error())
		return c.JSON(401, res)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokens
	cookie.Expires = time.Now().Add(time.Hour * 3)
	c.SetCookie(cookie)

	res := helper.ResponseJSON("Success", 200, "OK", tokens)
	return c.JSON(http.StatusOK, res)
}

func (h *handler) Logout(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		res := helper.ResponseJSON("Faled logout", 500, "error", err.Error())
		return c.JSON(500, res)
	}

	cookie.Value = ""
	cookie.Expires = time.Now()
	c.SetCookie(cookie)

	res := helper.ResponseJSON("Success Logout", 200, "OK", nil)
	return c.JSON(http.StatusOK, res)
}
