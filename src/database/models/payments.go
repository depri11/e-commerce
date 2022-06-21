package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderID         string             `json:"order_id" bson:"order_id"`
	TransactionId   string             `json:"transaction_id" bson:"transaction_id"`
	User            string             `json:"user" bson:"user"`
	Amount          string             `json:"amount" bson:"amount"`
	PaymentType     string             `json:"payment_type" bson:"payment_type"`
	Status          string             `json:"status" bson:"status"`
	TransactionTime string             `json:"transaction_time" bson:"transaction_time"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}
