package config

import (
	"context"
	"log"
	"os"
	"time"

	// "gorilla/config"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_NAME = "nontachai_test"

func MongoDbConnection() (*mongo.Client, context.Context, error) {
	godotenv.Load(".env")
	mongodbUri := os.Getenv("MONGO_DB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	return client, ctx, err
}
