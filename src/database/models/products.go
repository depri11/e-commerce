package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name"`
	Description    string             `json:"description,omitempty" bson:"description"`
	Specifications []Specification    `json:"specifications,omitempty" bson:"specifications"`
	Price          int32              `json:"price,omitempty" bson:"price"`
	CuttedPrice    int32              `json:"cutted_price,omitempty" bson:"cutted_price"`
	Images         []ProductImage     `json:"images,omitempty" bson:"images"`
	Brand          Brand              `json:"brand,omitempty" bson:"brand"`
	Category       string             `json:"category,omitempty" bson:"category"`
	Stock          int                `json:"stock,omitempty" bson:"stock"`
	Warranty       int                `json:"warranty,omitempty" bson:"warranty"`
	Ratings        float32            `json:"ratings,omitempty" bson:"ratings"`
	NumOfReviews   int                `json:"num_of_reviews,omitempty" bson:"num_of_reviews"`
	Reviews        []Review           `json:"reviews,omitempty" bson:"reviews"`
	User           []User             `json:"user,omitempty" bson:"user"`
	CreatedAt      time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

type Specification struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title"`
	Description string             `json:"description,omitempty" bson:"description"`
}

type ProductImage struct {
	ID  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Url string             `json:"url,omitempty" bson:"url"`
}

type Brand struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name"`
	Logo Logo               `json:"logo,omitempty" bson:"logo"`
}

type Logo struct {
	ID  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Url string             `json:"url,omitempty" bson:"url"`
}

type Review struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    string             `json:"user_id,omitempty" bson:"user_id"`
	ProductID string             `json:"product_id,omitempty" bson:"product_id"`
	Name      string             `json:"name,omitempty" bson:"name"`
	Rating    float32            `json:"rating,omitempty" bson:"rating"`
	Comment   string             `json:"comment,omitempty" bson:"comment"`
}
