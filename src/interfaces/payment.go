package interfaces

import "github.com/depri11/e-commerce/src/database/models"

type PaymentService interface {
	GetPaymentURL(transaction models.Payment, user models.User) (string, error)
}
