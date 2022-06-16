package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name"`
	Description    string             `json:"description" bson:"description"`
	Specifications []Specification    `json:"specifications" bson:"specifications"`
	Price          int32              `json:"price" bson:"price"`
	CuttedPrice    int32              `json:"cutted_price" bson:"cutted_price"`
	Images         []ProductImage     `json:"images" bson:"images"`
	Brand          Brand              `json:"brand" bson:"brand"`
	Category       string             `json:"category" bson:"category"`
	Stock          int                `json:"stock" bson:"stock"`
	Warranty       int                `json:"warranty" bson:"warranty"`
	Ratings        float32            `json:"ratings" bson:"ratings"`
	NumOfReviews   int                `json:"num_of_reviews" bson:"num_of_reviews"`
	Reviews        []Review           `json:"reviews" bson:"reviews"`
	User           []User             `json:"user" bson:"user"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
}

type Specification struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
}

type ProductImage struct {
	ID  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url string             `json:"url" bson:"url"`
}

type Brand struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Logo Logo               `json:"logo" bson:"logo"`
}

type Logo struct {
	ID  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url string             `json:"url" bson:"url"`
}

type Review struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID  string             `json:"user_id" bson:"user_id"`
	Name    string             `json:"name" bson:"name"`
	Rating  float32            `json:"rating" bson:"rating"`
	Comment string             `json:"comment" bson:"comment"`
}
