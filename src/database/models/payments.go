package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Amount int                `json:"amount" bson:"amount"`
}
