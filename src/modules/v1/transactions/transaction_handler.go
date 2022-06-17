package transaction

import (
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service interfaces.TransactionService
}

func NewHandler(service interfaces.TransactionService) *handler {
	return &handler{service}
}

func (h *handler) GetTransactions(c echo.Context) error {
	data, err := h.service.GetAll()
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

	data, err := h.service.ProcessPayment(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
