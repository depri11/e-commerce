package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/input"
)

type PaymentService interface {
	GetPaymentURL(orderID string, order *models.Order, user *input.RespUser) (string, error)
	ProcessPayment(input *input.OrderNotificationInput) (*helper.Res, error)
}
