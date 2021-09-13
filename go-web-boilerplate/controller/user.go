package controller

import (
	"net/http"
	"pikachu/model"
	"pikachu/repository"
	"pikachu/service"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// User ...
type User struct {
	userSvc  service.UserService
	userRepo repository.UserRepository
}

// NewUserController ...
func NewUserController(userSvc service.UserService, userRepo repository.UserRepository) UserController {
	return &User{
		userSvc:  userSvc,
		userRepo: userRepo,
	}
}

// NewUser ...
func (u *User) NewUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	zlog.With(ctx).Infow("[New request]")

	user := &model.User{}
	if err := c.Bind(user); err != nil {
		zlog.With(ctx).Warnw("Bind error", "user", user, "err", err)
		return response(c, http.StatusBadRequest, "Bind error")
	} else if !user.ValidateNewUser() {
		zlog.With(ctx).Warnw("NewUser ValidateNewUser failed", "user", user)
		return response(c, http.StatusBadRequest, "Validate failed")
	}

	if user, err = u.userSvc.NewUser(ctx, user); err != nil {
		zlog.With(ctx).Errorw("UserSvc NewUser failed", "user", user, "err", err)
		return response(c, http.StatusInternalServerError, "NewUser failed")
	}

	return response(c, http.StatusOK, "New Deal OK", user)
}

// GetUser ...
func (u *User) GetUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	uid := c.Param("uid")
	zlog.With(ctx).Infow("[New request]", "uid", uid)

	if _, err = uuid.Parse(uid); err != nil {
		zlog.With(ctx).Warnw("ID is not valid", "uid", uid, "err", err)
		return response(c, http.StatusBadRequest, "User is not valid")
	}

	user := &model.User{}
	if user, err = u.userSvc.GetUser(ctx, uid); err != nil {
		zlog.With(ctx).Warnw("UserSvc GetUser failed", "uid", uid, "err", err)
		return response(c, http.StatusInternalServerError, "GetUser failed")
	}

	return response(c, http.StatusOK, "GetUser OK", user)
}

// UpdateUser ...
func (u *User) UpdateUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	zlog.With(ctx).Infow("[New request]")

	user := &model.User{}
	if err := c.Bind(user); err != nil {
		zlog.With(ctx).Warnw("Bind error", "user", user, "err", err)
		return response(c, http.StatusBadRequest, "Bind error")
	}
	if user, err = u.userSvc.UpdateUser(ctx, uid, user); err != nil {
		zlog.With(ctx).Errorw("UserSvc NewUser failed", "user", user, "err", err)
		return response(c, http.StatusInternalServerError, "UpdateUser failed")
	}

	return response(c, http.StatusOK, "Update Deal OK", user)
}

// DeleteUser ...
func (u *User) DeleteUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	uid := c.Param("uid")
	zlog.With(ctx).Infow("[New request]", "uid", uid)

	if _, err = uuid.Parse(uid); err != nil {
		zlog.With(ctx).Warnw("ID is not valid", "uid", uid, "err", err)
		return response(c, http.StatusBadRequest, "User is not valid")
	}

	if err = u.userSvc.DeleteUser(ctx, uid); err != nil {
		zlog.With(ctx).Errorw("UserSvc DeleteUser failed", "uid", uid, "err", err)
		return response(c, http.StatusInternalServerError, "DeleteUser failed")
	}

	return response(c, http.StatusOK, "Delete User OK")
}
