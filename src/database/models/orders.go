package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	ShippingInfo ShippingInfo       `json:"shipping_info" bson:"shippingInfo"`
	Items        []Item             `json:"items" bson:"items"`
	User         User               `json:"user" bson:"user"`
	PaymentInfo  PaymentInfo        `json:"payment_info" bson:"paymentInfo"`
	// PaidAt       time.Time          `json:"paid_at" bson:"paidAt"`
	TotalPrice float64 `json:"total_price" bson:"totalPrice"`
	Status     string  `json:"status" bson:"status"`
	// DeliveredAt  time.Time          `json:"delivered_at" bson:"deliveredAt"`
	// ShippedAt    time.Time          `json:"shipped_at" bson:"shippedAt"`
	CreatedAt time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" bson:"updatedAt"`
}

type PaymentInfo struct {
	PaymentID string `json:"payment_id" bson:"paymentID"`
	Status    string `json:"status" bson:"status"`
}

type Item struct {
	Name     string `json:"name" bson:"name"`
	Price    int    `json:"price" bson:"price"`
	Quantity int    `json:"quantity" bson:"quantity"`
	// Image    string  `json:"image" bson:"image"`
	Product Product `json:"product" bson:"product"`
}

type ShippingInfo struct {
	Address string `bson:"address"`
	City    string `bson:"city"`
	State   string `bson:"state"`
	Country string `bson:"country"`
	Pincode string `bson:"pincode"`
	Phone   string `bson:"phone"`
}
