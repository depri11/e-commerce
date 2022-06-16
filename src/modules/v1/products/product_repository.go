package products

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

func (r *repository) FindAll() ([]models.Product, error) {
	ctx := context.TODO()

	var products []models.Product

	cursor, _ := r.C.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) FindByID(id string) (*models.Product, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	product := &models.Product{}

	err = r.C.FindOne(ctx, bson.M{"_id": p}).Decode(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *repository) Insert(user *models.Product) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()
	return r.C.InsertOne(ctx, user)
}

func (r *repository) Update(id string, product *models.Product) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	return r.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": product})
}

func (r *repository) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	return r.C.DeleteOne(ctx, bson.M{"_id": p})
}

func (r *repository) Search(query string) ([]models.Product, error) {
	ctx := context.TODO()

	var products []models.Product

	filter := bson.M{
		"$or": []bson.M{
			{
				"name": bson.M{
					"$regex": primitive.Regex{
						Pattern: query,
						Options: "i",
					},
				},
			},
			{
				"description": bson.M{
					"$regex": primitive.Regex{
						Pattern: query,
						Options: "i",
					},
				},
			},
		},
	}

	cursor, err := r.C.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	return products, nil
}
