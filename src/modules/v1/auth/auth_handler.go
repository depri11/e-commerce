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

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) SigIn(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	res := h.service.Login(user)
	return c.JSON(http.StatusOK, res.Data)
}
