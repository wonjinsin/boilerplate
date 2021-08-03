package service

import (
	"pikachu/repository"

	"pikachu/model"
)

// Service ...
type Service struct {
	User UserService
}

// Init ...
func Init(repo *repository.Repository) (*Service, error) {
	userSvc := NewUserService(repo.User)

	return &Service{User: userSvc}, nil
}

// UserService ...
type UserService interface {
	NewUser() *model.User
}
