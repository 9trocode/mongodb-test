package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Define the MongoDB connection URI (adjust if necessary)
	uri := "mongodb://pipeops:624a8aeb7d63fdd46df297d38@lingering-sun.big-action-beta.svc.pipeops.internal:27017"

	// Create a new client and set the options
	clientOptions := options.Client().ApplyURI(uri)

	// Create a context with a timeout to avoid hanging connections
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	time.Sleep(30 * time.Minute)

	// Disconnect from MongoDB when done
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
	}
}
