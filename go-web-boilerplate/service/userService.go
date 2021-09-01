package service

import (
	"context"
	"pikachu/model"
	"pikachu/repository"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserService ...
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userUsecase{
		userRepo: userRepo,
	}
}

// NewUser ...
func (u *userUsecase) NewUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	zlog.With(ctx).Infow("[New Service Request]", "user", user)
	if ruser, err = u.userRepo.NewUser(ctx, user); err != nil {
		zlog.With(ctx).Errorw("UserRepo NewUser Failed", "user", user)
		return nil, err
	}

	return ruser, nil
}

// GetUser ...
func (u *userUsecase) GetUser(ctx context.Context, uid string) (ruser *model.User, err error) {
	zlog.With(ctx).Infow("[New Service Request]", "uid", uid)
	if ruser, err = u.userRepo.GetUser(ctx, uid); err != nil {
		zlog.With(ctx).Errorw("UserRepo GetUser Failed", "uid", uid, "err", err)
		return nil, err
	}

	return ruser, nil
}

// UpdateUser ...
func (u *userUsecase) UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	zlog.With(ctx).Infow("[New Service Request]", "user", user)
	if ruser, err = u.userRepo.UpdateUser(ctx, user); err != nil {
		zlog.With(ctx).Errorw("UserRepo UpdateUser Failed", "user", user)
		return nil, err
	}

	return ruser, nil
}

// DeleteUser ...
func (u *userUsecase) DeleteUser(ctx context.Context, uid string) (err error) {
	zlog.With(ctx).Infow("[New Service Request]", "uid", uid)
	if err = u.userRepo.DeleteUser(ctx, uid); err != nil {
		zlog.With(ctx).Errorw("UserRepo DeleteUser Failed", "uid", uid, "err", err)
		return nil, err
	}

	return err, nil
}
