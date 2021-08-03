package controller

import (
	"pikachu/repository"
	"pikachu/service"

	"github.com/labstack/echo"
)

type user struct {
	userSvc  service.UserService
	userRepo repository.UserRepository
}

// NewUserController ...
func NewUserController(userSvc service.UserService, userRepo repository.UserRepository) *user {
	return &user{
		userSvc:  userSvc,
		userRepo: userRepo,
	}
}

func (u *user) GetUser(c echo.Context) (err error) {
	return nil
}
