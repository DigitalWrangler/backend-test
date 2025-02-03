package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"my-go-project/internal/api"
	"my-go-project/internal/service"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	const database = "userdb"

	// Set up MongoDB connection with context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get MongoDB URI from environment variable or use default with authentication
	mongoURI := getEnv("MONGODB_URI", "mongodb://root:gotest@localhost:27017/userdb?authSource=admin")

	// Set up client options for MongoDB connection
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create MongoDB client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		// Disconnect MongoDB client on application shutdown
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB:", err)
		}
	}()

	// Wait for MongoDB to be ready and log connection retries
	if err := waitForMongo(client, ctx); err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	// Initialize UserService with the connected MongoDB client
	userService := service.NewUserService(client.Database(database))

	// Initialize Echo web framework for the server
	e := echo.New()

	// ROUTES
	e.GET("/users", api.GetUsers(userService))
	e.POST("/users", api.CreateUser(userService))
	e.PUT("/users/:id/deactivate", api.DeactivateUser(userService))

	// Start the server on the defined port
	port := getEnv("PORT", "8888")
	log.Printf("Server starting on port %s...", port)
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Helper function to get environment variables with a fallback option
func getEnv(key, fallback string) string {
	// Check if the environment variable exists and return its value, otherwise return the fallback
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Helper function to check MongoDB connection readiness
func waitForMongo(client *mongo.Client, ctx context.Context) error {
	retries := 0
	maxRetries := 30

	// Try to ping MongoDB server up to maxRetries times to ensure it's ready
	for retries < maxRetries {
		err := client.Ping(ctx, nil) // Ping the MongoDB server to check connectivity
		if err == nil {
			log.Println("Successfully connected to MongoDB!")
			return nil
		}

		retries++
		log.Printf("MongoDB not ready yet, retrying (%d/%d)... Error: %v", retries, maxRetries, err)
		time.Sleep(time.Second) // Wait before retrying
	}

	// Return an error if MongoDB connection could not be established after max retries
	return fmt.Errorf("could not connect to MongoDB after %d retries", maxRetries)
}
