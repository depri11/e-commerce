package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
)

type PaymentService interface {
	GetPaymentURL(transaction *models.Transaction, user *models.User) (string, error)
	ProcessPayment(input *models.TransactionNotification) (*helper.Res, error)
}
