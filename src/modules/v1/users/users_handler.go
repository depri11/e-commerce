package users

import (
	"fmt"
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service interfaces.UserService
}

func NewHandler(service interfaces.UserService) *handler {
	return &handler{service}
}

func (h *handler) FindAll(c echo.Context) error {
	data, err := h.service.FindAll()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	res, err := h.service.Insert(&user)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
