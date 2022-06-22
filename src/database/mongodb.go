package database

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB() (*mongo.Database, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	clientOption := options.Client().ApplyURI("mongodb://localhost:27017/")

	db, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = db.Connect(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")

	return db.Database("e-commerce"), nil

}
