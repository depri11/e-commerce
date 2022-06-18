package orders

import (
	"context"

	"github.com/depri11/e-commerce/src/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	C *mongo.Collection
}

func NewRepository(C *mongo.Collection) *repository {
	return &repository{C}
}

func (r *repository) FindAll() ([]*models.Order, error) {
	ctx := context.TODO()
	var orders []*models.Order

	cursor, err := r.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *repository) Insert(order models.Order) (*models.Order, error) {
	ctx := context.TODO()
	_, err := r.C.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *repository) Delete(id string) (*mongo.DeleteResult, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	return r.C.DeleteOne(ctx, bson.M{"_id": oid})
}

// func (r *repository) FindOrderById(id string) (*models.Order, error) {
// 	// oid := primitive.ObjectID.Hex(id)
// 	ctx := context.TODO()
// 	var order *models.Order

// 	err := r.C.FindOne(ctx, bson.M{"_id": oid}).Decode(&order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return order, nil
// }

// func (r *repository) FindByUserId(id string) ([]*models.Order, error) {
// 	ctx := context.TODO()
// 	var order *models.Order

// 	cursor, err := r.C.FindOne(ctx, bson.M{""})

// 	return orders, nil
// }
