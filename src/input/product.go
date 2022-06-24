package input

import (
	"time"

	"github.com/depri11/e-commerce/src/database/models"
)

type CreateProductInput struct {
	Name           string                 `json:"name,omitempty" validate:"required"`
	Description    string                 `json:"description,omitempty" validate:"required"`
	Specifications []models.Specification `json:"specifications,omitempty" validate:"required"`
	Price          int32                  `json:"price,omitempty" validate:"required"`
	CuttedPrice    int32                  `json:"cutted_price,omitempty"`
	Images         []models.Image         `json:"images"`
	Brand          models.Brand           `json:"brand,omitempty"`
	Category       string                 `json:"category,omitempty" validate:"required"`
	Stock          int                    `json:"stock,omitempty"`
	Warranty       int                    `json:"warranty,omitempty"`
}

type CreateReviewInput struct {
	Comment string  `json:"comment,omitempty"`
	Rating  float64 `json:"rating,omitempty" validate:"required"`
}

type UpdateProductInput struct {
	Name           string                 `json:"name,omitempty" validate:"required"`
	Description    string                 `json:"description,omitempty" validate:"required"`
	Specifications []models.Specification `json:"specifications,omitempty" validate:"required"`
	Price          int32                  `json:"price,omitempty" validate:"required"`
	CuttedPrice    int32                  `json:"cutted_price,omitempty"`
	Images         []models.Image         `json:"images"`
	Brand          models.Brand           `json:"brand,omitempty"`
	Category       string                 `json:"category,omitempty" validate:"required"`
	Stock          int                    `json:"stock,omitempty"`
	Warranty       int                    `json:"warranty,omitempty"`
	UpdatedAt      time.Time              `json:"updated_at,omitempty"`
}
