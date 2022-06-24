package input

import "github.com/depri11/e-commerce/src/database/models"

type CreateOrderInput struct {
	ShippingInfo models.ShippingInfo `json:"shipping_info" validate:"required"`
	Products     []string            `json:"items" validate:"required"`
	TotalPrice   float64             `json:"total_price" validate:"required"`
}

type OrderNotificationInput struct {
	OrderStatus string `json:"transaction_status" validate:"required"`
	OrderID     string `json:"order_id" validate:"required"`
	PaymentType string `json:"payment_type" validate:"required"`
	FraudStatus string `json:"fraud_status" validate:"required"`
}
