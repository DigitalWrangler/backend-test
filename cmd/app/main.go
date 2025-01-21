package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"log/slog"
	"my-go-project/internal/api"
	"my-go-project/internal/service"
	"os"
	"time"
)

func main() {
	const database = "userdb"

	// Set up MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get MongoDB URI from environment variable or use default
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			slog.Error("Failed to disconnect from MongoDB", "error", err)
		}
	}(client, ctx)

	// Wait for MongoDB to be ready
	err = waitForMongo(client, ctx)
	if err != nil {
		panic(err)
	}

	userService := service.NewUserService(client.Database(database))

	// Initialize echo
	e := echo.New()
	// Setup API routes
	e.GET("/users", api.GetUsers(userService))
	// TODO add more routes

	// Start server
	port := getEnv("PORT", "8888")
	slog.Info("Server starting...", "port", port)
	log.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

// Helper function to get environment variables
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Helper function to wait for MongoDB to be ready
func waitForMongo(client *mongo.Client, ctx context.Context) error {
	retries := 0
	maxRetries := 30

	for retries < maxRetries {
		err := client.Ping(ctx, nil)
		if err == nil {
			log.Println("Successfully connected to MongoDB!")
			return nil
		}

		retries++
		log.Printf("Waiting for MongoDB... (%d/%d)\n", retries, maxRetries)
		time.Sleep(time.Second)
	}

	return fmt.Errorf("could not connect to MongoDB after %d retries", maxRetries)
}
