package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"my-go-project/internal/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func main() {
	// Set up MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Get MongoDB URI from environment variable or use default
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Wait for MongoDB to be ready
	err = waitForMongo(client, ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize collection
	collection = client.Database("assetsdb").Collection("assets")

	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/assets", getAssets).Methods("GET")

	// Start server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getAssets(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var assets []models.Asset
	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &assets); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assets)
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