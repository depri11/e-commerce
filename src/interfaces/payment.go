package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
)

type PaymentService interface {
	GetPaymentURL(orderID string, order *models.Order, user *models.User) (string, error)
	ProcessPayment(input *models.OrderNotification) (*helper.Res, error)
}
