package database

import (
	"context"
	"fmt"
	"os"
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

	mongo_uri := os.Getenv("MONGODB_URI")
	name := os.Getenv("MONGODB_NAME")

	clientOption := options.Client().ApplyURI(mongo_uri)

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

	return db.Database(name), nil

}
