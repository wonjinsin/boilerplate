package service

import (
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

func (u *userUsecase) NewUser(user *model.User) *model.User {
	return user
}
