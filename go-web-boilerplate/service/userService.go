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
	zlog.With(ctx).Infow("Service NewUser", "user", user)
	if ruser, err = u.userRepo.NewUser(ctx, user); err != nil {
		zlog.With(ctx).Errorw("NewUser Failed", user)
		return nil, err
	}

	return ruser, nil
}
