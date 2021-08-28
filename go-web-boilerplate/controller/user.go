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
		zlog.With(ctx).Warnw("Bind Error", "user", user, "err", err)
		return response(c, http.StatusBadRequest, "Bind Error")
	}

	if user, err = u.userSvc.NewUser(ctx, user); err != nil {
		zlog.With(ctx).Errorw("UserSvc NewUser Failed", "user", user, "err", err)
		return response(c, http.StatusInternalServerError, "NewUser Failed", err)
	}

	return response(c, http.StatusOK, "New Deal OK", user)
}
