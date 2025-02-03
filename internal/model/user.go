package models

import "time"

// User represents a user in MongoDB
type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`      // MongoDB _id
	Name      string    `json:"name" bson:"name"`             // User's name
	Email     string    `json:"email" bson:"email"`           // User's email address
	BirthDate string    `json:"birth_date" bson:"birth_date"` // Birth date in the format YYYY-MM-DD
	City      string    `json:"city" bson:"city"`             // City of the user
	Age       int       `json:"age,omitempty"`                // Age calculated dynamically, optional in JSON response
	Active    bool      `json:"active" bson:"active"`         // User's active status
	CreatedAt time.Time `json:"created_at" bson:"created_at"` // Timestamp when user was created
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` // Timestamp when user was last updated
}

// CalculateAge calculates the age based on the BirthDate
func (u *User) CalculateAge() int {
	birthDate, err := time.Parse("2006-01-02", u.BirthDate) // Parse the birthdate string into a time.Time object
	if err != nil {
		return 0 // Return 0 if the birthdate cannot be parsed
	}
	now := time.Now()
	years := now.Year() - birthDate.Year()   // Calculate the age in years
	if now.YearDay() < birthDate.YearDay() { // Adjust if birthday hasn't occurred yet this year
		years--
	}
	return years
}
