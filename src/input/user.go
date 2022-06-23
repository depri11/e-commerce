package input

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInput struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Gender   string `json:"gender,omitempty" validate:"required"`
	Avatar   string `json:"avatar,omitempty"`
	Password string `json:"password,omitempty" validate:"required,min=6"`
}

type RespUser struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string             `json:"name,omitempty"`
	Email  string             `json:"email,omitempty"`
	Gender string             `json:"gender,omitempty"`
	Avatar string             `json:"avatar,omitempty"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" validate:"required,email"`
}
