package transaction

import (
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service        interfaces.TransactionService
	paymentService interfaces.PaymentService
}

func NewHandler(service interfaces.TransactionService, paymentService interfaces.PaymentService) *handler {
	return &handler{service, paymentService}
}

func (h *handler) GetProductTransactions(c echo.Context) error {
	id := c.Param("id")

	data, err := h.service.GetByProductID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetUserTransactions(c echo.Context) error {
	id := c.Request().Header.Get("user_id")

	data, err := h.service.GetByUserID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) CreateTransaction(c echo.Context) error {
	var transaction models.Transaction

	id := c.Request().Header.Get("user_id")

	if err := c.Bind(&transaction); err != nil {
		return err
	}

	data, err := h.service.Create(id, &transaction)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetNotification(c echo.Context) error {
	var input models.TransactionNotification

	if err := c.Bind(&input); err != nil {
		return err
	}

	data, err := h.paymentService.ProcessPayment(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
