package service

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"my-go-project/internal/model"
)

// This interface is just a suggestion, feel free to change anything!

type UserService interface {
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) error
	DeactivateUser(id string) error
}

type userService struct {
}

func (s *userService) GetUsers() ([]model.User, error) {
	// TODO implement
	return nil, nil
}

func (s *userService) CreateUser(user model.User) error {
	// TODO implement
	return nil
}

func (s *userService) DeactivateUser(id string) error {
	// TODO implement
	return nil
}

func NewUserService(database *mongo.Database) UserService {
	return &userService{}
}
