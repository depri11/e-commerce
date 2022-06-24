package orders

import (
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/input"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service        interfaces.OrderService
	paymentService interfaces.PaymentService
}

func NewHandler(service interfaces.OrderService, paymentService interfaces.PaymentService) *handler {
	return &handler{service, paymentService}
}

func (h *handler) GetAllOrders(c echo.Context) error {
	data, err := h.service.GetAllOrders()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetOrderDetails(c echo.Context) error {
	id := c.Param("id")

	data, err := h.service.FindByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) MyOrders(c echo.Context) error {
	id := c.Request().Header.Get("user_id")

	data, err := h.service.FindByUserID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) NewOrder(c echo.Context) error {
	id := c.Request().Header.Get("user_id")

	var input input.CreateOrderInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(500, err.Error())
	}

	if err := helper.ValidationError(input); err != nil {
		return c.JSON(400, err.Error())
	}

	data, err := h.service.Create(id, &input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) UpdateOrder(c echo.Context) error {
	var order models.Order
	id := c.Param("id")

	if err := c.Bind(&order); err != nil {
		return err
	}

	data, err := h.service.Update(id, &order)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) DeleteOrder(c echo.Context) error {
	id := c.Param("id")

	data, err := h.service.Delele(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetNotificationOrder(c echo.Context) error {
	var input input.OrderNotificationInput

	if err := c.Bind(&input); err != nil {
		return err
	}

	data, err := h.paymentService.ProcessPayment(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
