package orders

import (
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service interfaces.OrderService
}

func NewHandler(service interfaces.OrderService) *handler {
	return &handler{service}
}

func (h *handler) GetAllOrders(c echo.Context) error {
	data, err := h.service.GetAllOrders()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) Create(c echo.Context) error {
	var order models.Order

	if err := c.Bind(&order); err != nil {
		return err
	}

	data, err := h.service.Create(order)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) Delete(c echo.Context) error {
	id := c.Param("id")

	data, err := h.service.Delele(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
