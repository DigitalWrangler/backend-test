package api

import (
	models "my-go-project/internal/model"
	"my-go-project/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUsers is the handler for fetching all users with sorting
func GetUsers(userService *service.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the sorting parameters from the query string
		sortBy := c.QueryParam("sortBy") // e.g., "name", "email", or "age"
		order := c.QueryParam("order")   // e.g., "asc" or "desc"

		// Validate the sort order to be either "asc" or "desc"
		if order != "asc" && order != "desc" && order != "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid order. Use 'asc' or 'desc'.",
			})
		}

		// Fetch all users with sorting
		users, err := userService.GetAllUsers(sortBy, order)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error fetching users",
			})
		}

		// If no users are found, return a 404 with a message
		if len(users) == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "No users found",
			})
		}

		// Return the users in a 200 OK response
		return c.JSON(http.StatusOK, users)
	}
}

// CreateUser is the handler for creating a new user
func CreateUser(userService *service.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind the incoming JSON request body to the User model
		var newUser models.User
		if err := c.Bind(&newUser); err != nil {
			// If binding fails, return a 400 Bad Request
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid user data",
			})
		}

		// Call the service to create the new user
		err := userService.CreateUser(&newUser)
		if err != nil {
			// If an error occurs while creating the user, return a 500 Internal Server Error
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error creating user",
			})
		}

		// Return the newly created user with a 201 Created status
		return c.JSON(http.StatusCreated, newUser)
	}
}

// DeactivateUser is the handler for deactivating a user
func DeactivateUser(userService *service.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("id")

		// Call the service to deactivate the user
		err := userService.DeactivateUser(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error deactivating user",
			})
		}

		// Respond with success message
		return c.JSON(http.StatusOK, map[string]string{
			"message": "User deactivated successfully",
		})
	}
}
