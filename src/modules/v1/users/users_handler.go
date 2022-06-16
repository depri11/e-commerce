package users

import (
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
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetUserID(c echo.Context) error {
	id := c.Param("id")
	data, err := h.service.GetUserID(id)
	if err != nil {
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

func (h *handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	res, err := h.service.Update(id, &user)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)

}

func (h *handler) DeletUser(c echo.Context) error {
	id := c.Param("id")
	res, err := h.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}