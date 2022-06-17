package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID  string             `json:"product_id" bson:"product_id"`
	UserID     string             `json:"user_id" bson:"user_id"`
	Amount     int                `json:"amount" bson:"amount"`
	Status     string             `json:"status" bson:"status"`
	Code       string             `json:"code" bson:"code"`
	PaymentURL string             `json:"payment_url" bson:"payment_url"`
	User       User               `json:"user" bson:"user"`
	Product    Product            `json:"product" bson:"product"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type TransactionNotification struct {
	TransactionStatus string `json:"status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
