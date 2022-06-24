package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	OrderID      string             `json:"order_id" bson:"order_id"`
	ShippingInfo ShippingInfo       `json:"shipping_info" bson:"shippingInfo"`
	Products     []string           `json:"items" bson:"items"`
	UserID       string             `json:"user" bson:"user"`
	TotalPrice   float64            `json:"total_price" bson:"totalPrice"`
	Status       string             `json:"status" bson:"status"`
	PaymentURL   string             `json:"payment_url" bson:"paymentUrl"`
	PaidAt       time.Time          `json:"paid_at" bson:"paidAt"`
	DeliveredAt  *time.Time         `json:"delivered_at,omitempty" bson:"deliveredAt,omitempty"`
	ShippedAt    *time.Time         `json:"shipped_at,omitempty" bson:"shippedAt,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updatedAt"`
}

type ShippingInfo struct {
	Address string `json:"address" bson:"address"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Pincode string `json:"pincode" bson:"pincode"`
	Phone   string `json:"phone" bson:"phone"`
}
