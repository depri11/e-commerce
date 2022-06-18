package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	ShippingInfo ShippingInfo       `json:"shipping_info" bson:"shippingInfo"`
	Items        []string           `json:"items" bson:"items"`
	UserID       string             `json:"user" bson:"user"`
	PaymentInfo  PaymentInfo        `json:"payment_info" bson:"paymentInfo"`
	TotalPrice   float64            `json:"total_price" bson:"totalPrice"`
	Status       string             `json:"status" bson:"status"`
	PaidAt       time.Time          `json:"paid_at" bson:"paidAt"`
	DeliveredAt  time.Time          `json:"delivered_at" bson:"deliveredAt"`
	ShippedAt    time.Time          `json:"shipped_at" bson:"shippedAt"`
	CreatedAt    time.Time          `json:"created_at" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updatedAt"`
}

type PaymentInfo struct {
	PaymentID string `json:"payment_id" bson:"paymentID"`
	Status    string `json:"status" bson:"status"`
}

type ShippingInfo struct {
	Address string `json:"address" bson:"address"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Pincode string `json:"pincode" bson:"pincode"`
	Phone   string `json:"phone" bson:"phone"`
}
