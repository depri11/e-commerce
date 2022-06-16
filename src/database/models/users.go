package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name"`
	Email    string             `json:"email,omitempty" bson:"email"`
	Gender   string             `json:"gender,omitempty" bson:"gender"`
	Password string             `json:"password,omitempty" bson:"password"`
	// Avatar          Avatar    `json:"avatar" bson:"avatar,omitempty"`
	Role      string    `json:"role,omitempty" bson:"role"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	// ResetPassToken  string    `json:"reset_pass_token" bson:"reset_pass_token,omitempty"`
	// ResetPassExpire time.Time `json:"reset_pass_expire" bson:"reset_pass_expire,omitempty"`
}

type Avatar struct {
	ID  primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Url string             `json:"url,omitempty" bson:"url"`
}
