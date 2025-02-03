package service

import (
	"context"
	"log"
	models "my-go-project/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserService provides methods to interact with the User data in MongoDB
type UserService struct {
	collection *mongo.Collection
}

// NewUserService creates a new UserService
func NewUserService(db *mongo.Database) *UserService {
	// Ensure we're using the correct collection name (users)
	return &UserService{
		collection: db.Collection("users"), // Using "users" collection
	}
}

// GetAllUsers fetches all users from the database and calculates their age, with sorting support
func (s *UserService) GetAllUsers(sortBy, order string) ([]models.User, error) {
	var users []models.User

	//Sort options
	sortOptions := bson.D{}
	if sortBy != "" {
		// If sorting by 'age', sort by 'birth_date' instead
		if sortBy == "age" {
			sortBy = "birth_date"
		}

		// Determine the sort order (ascending or descending)
		orderInt := 1 // Default to ascending
		if order == "desc" {
			orderInt = -1 // Descending order
		}

		// Apply sorting by the specified field
		sortOptions = bson.D{
			{Key: sortBy, Value: orderInt},
		}
	}

	// Fetch all users with the specified sort options
	cursor, err := s.collection.Find(context.Background(), bson.D{}, &options.FindOptions{
		Sort: sortOptions,
	})
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each user
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user:", err)
			continue
		}

		// Calculate age dynamically
		user.Age = user.CalculateAge()

		// Append the user to the list
		users = append(users, user)
	}

	// Check if there was an error while iterating through the cursor
	if err := cursor.Err(); err != nil {
		log.Println("Error iterating through users:", err)
		return nil, err
	}

	// Return the list of users
	return users, nil
}

// CreateUser inserts a new user into the database
func (s *UserService) CreateUser(newUser *models.User) error {
	// Set the timestamps for creation and update
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	// Insert the user into the database
	_, err := s.collection.InsertOne(context.Background(), newUser)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	log.Printf("User created: %+v\n", newUser)
	return nil
}

// DeactivateUser updates the active status of a user
func (s *UserService) DeactivateUser(userID string) error {
	// Convert userID from string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Error converting user ID to ObjectID:", err)
		return err
	}

	// Update the active field of the user with the provided ObjectID
	_, err = s.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},                 // Find user by ObjectID
		bson.M{"$set": bson.M{"active": false}}, // Set active field to false
	)
	if err != nil {
		log.Println("Error deactivating user:", err)
		return err
	}

	return nil
}
