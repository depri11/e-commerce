package users

import (
	"context"
	"fmt"

	"github.com/depri11/e-commerce/src/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *repository) FindByID(id string) (*models.User, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	user := &models.User{}

	err = r.C.FindOne(ctx, bson.M{"_id": p}).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Insert(user *models.User) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()
	return r.C.InsertOne(ctx, user)
}

func (r *repository) Update(id string, user *models.User) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	return r.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": user})
}

func (r *repository) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	return r.C.DeleteOne(ctx, bson.M{"_id": p})
}
