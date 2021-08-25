package controller

import (
	"net/http"
	"pikachu/model"
	"pikachu/repository"
	"pikachu/service"

	"github.com/labstack/echo"
)

// User ...
type User struct {
	userSvc  service.UserService
	userRepo repository.UserRepository
}

// NewUserController ...
func NewUserController(userSvc service.UserService, userRepo repository.UserRepository) *User {
	return &User{
		userSvc:  userSvc,
		userRepo: userRepo,
	}
}

// NewUser ...
func (u *User) NewUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	zlog.With(ctx).Infow("[New Request]")

	user := &model.User{}
	if err := c.Bind(user); err != nil {
		zlog.With(ctx).Warnw("Bind Error", "err", err, "user", user)
		return response(c, http.StatusBadRequest, "Bind Error")
	}

	return nil
}
