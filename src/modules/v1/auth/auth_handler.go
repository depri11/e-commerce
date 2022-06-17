package auth

import (
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
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

	res, err := h.service.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
