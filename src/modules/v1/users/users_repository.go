package users

import (
	"context"
	"fmt"

	"github.com/depri11/e-commerce/src/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	C *mongo.Collection
}

func NewRepository(c *mongo.Collection) *repository {
	return &repository{C: c}
}

func (r *repository) FindAll() ([]*models.User, error) {
	ctx := context.TODO()
	cursor, err := r.C.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var users []*models.User
	if err := cursor.All(ctx, &users); err != nil {
		fmt.Println(err.Error())

		return nil, err
	}

	return users, nil
}

func (r *repository) Insert(user *models.User) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()
	return r.C.InsertOne(ctx, user)
}
