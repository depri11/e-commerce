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
	Ratings        float64            `json:"ratings" bson:"ratings"`
	NumOfReviews   int                `json:"num_of_reviews" bson:"num_of_reviews"`
	Reviews        []*Review          `json:"reviews,omitempty" bson:"reviews"`
	CreatedAt      time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

type Specification struct {
	Title       string `json:"title,omitempty" bson:"title"`
	Description string `json:"description,omitempty" bson:"description"`
}

type ProductImage struct {
	Url       string `json:"url,omitempty" bson:"url"`
	IsPrimary bool   `json:"is_primary,omitempty" bson:"is_primary"`
}

type Brand struct {
	Name string `json:"name,omitempty" bson:"name"`
	Logo string `json:"logo,omitempty" bson:"logo"`
}

type Review struct {
	UserID    string  `json:"user_id,omitempty" bson:"user_id"`
	ProductID string  `json:"product_id,omitempty" bson:"product_id"`
	Fullname  string  `json:"fullname,omitempty" bson:"fullname"`
	Rating    float64 `json:"rating,omitempty" bson:"rating"`
	Comment   string  `json:"comment,omitempty" bson:"comment"`
}

type ReviewInput struct {
	Review []*Review `json:"reviews,omitempty" bson:"reviews"`
}
