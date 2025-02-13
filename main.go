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
	// Adjust these values as needed:
	username := "pipeops_user"
	password := "5c9c49c9d42d744e337d04fea"
	host := "young-snowflake.big-action-beta.svc.pipeops.internal:27017"
	authSource := "admin" // Change if the user was created in a different database

	// Construct the connection URI. Try first without forcing an authMechanism.
	uri := fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s", username, password, host, authSource)
	// If needed, force SCRAM-SHA-256:
	// uri := fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s&authMechanism=SCRAM-SHA-256", username, password, host, authSource)
	fmt.Println("Connecting using URI:", uri)

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping MongoDB to verify the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	time.Sleep(30 * time.Minute)

	// Disconnect when done
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
	}
}
