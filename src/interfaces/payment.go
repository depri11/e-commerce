package interfaces

import "github.com/depri11/e-commerce/src/database/models"

type PaymentService interface {
	GetPaymentURL(transaction *models.Transaction, user *models.User) (string, error)
}
