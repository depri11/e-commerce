package auth

import (
	"net/http"
	"time"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
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
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	tokens, err := h.service.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
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
		return err
	}

	cookie.Value = ""
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	res := helper.ResponseJSON("Success Logout", 200, "OK", nil)
	return c.JSON(http.StatusOK, res)
}
