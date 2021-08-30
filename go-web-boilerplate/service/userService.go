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
		zlog.With(ctx).Errorw("NewUser Failed", user)
		return nil, err
	}

	return ruser, nil
}

// GetUser ...
func (u *userUsecase) GetUser(ctx context.Context, uid string) (ruser *model.User, err error) {
	zlog.With(ctx).Infow("[New Service Request]", "uid", uid)
	if ruser, err = u.userRepo.GetUser(ctx, uid); err != nil {
		zlog.With(ctx).Errorw("GetUser Failed", "uid", uid, "err", err)
		return nil, err
	}

	return ruser, nil
}
