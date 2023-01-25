package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DefaultDatabase = "blogdb"
	PostsCollection = "posts"
)

func Init() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectionURI := "mongodb://localhost:27017/"
	log.Println("Connecting to MongoDB")

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(connectionURI),
	)

	if err != nil {
		log.Fatalf("Unable to create Mongo client: %v\n", err)
	}

	// NOTE: The mongo.Client.Connect method doesn't ensure a database has been connected to before returning,
	// it just starts background monitoring threads that will attempt to discover a mongo deployment.
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	client.Database(DefaultDatabase).Collection(PostsCollection)

	log.Println("MongoDB connection successful")

	return client
}
