package models

type Payment struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Amount int    `json:"amount" bson:"amount"`
}
