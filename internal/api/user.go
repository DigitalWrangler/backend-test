package api

import (
	"github.com/labstack/echo/v4"
	"my-go-project/internal/service"
	"net/http"
)

func GetUsers(userService service.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := userService.GetUsers()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
}

func CreateUser(userService service.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO implement
		return c.NoContent(http.StatusOK)
	}
}

func DeleteUser(userService service.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO implement
		return c.NoContent(http.StatusOK)
	}
}
