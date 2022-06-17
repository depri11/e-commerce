package transaction

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

func NewRepository(c *mongo.Collection) *repository {
	return &repository{C: c}
}

func (r *repository) FindAll() ([]*models.Transaction, error) {
	ctx := context.TODO()
	var transactions []*models.Transaction

	cursor, err := r.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *repository) FindByID(id string) (*models.Transaction, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	transaction := &models.Transaction{}

	err = r.C.FindOne(ctx, bson.M{"_id": p}).Decode(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) FindByProductId(id string) (*models.Transaction, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	transaction := &models.Transaction{}

	err = r.C.FindOne(ctx, bson.M{"product_id": p}).Decode(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) FindByUserId(id string) (*models.Transaction, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	transaction := &models.Transaction{}

	err = r.C.FindOne(ctx, bson.M{"user_id": p}).Decode(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) Insert(transaction *models.Transaction) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()
	return r.C.InsertOne(ctx, transaction)
}

func (r *repository) Update(id string, user *models.Transaction) (*models.Transaction, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	transaction := &models.Transaction{}

	_, err = r.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
