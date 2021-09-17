package repository

import (
	"context"
	"pikachu/model"

	"github.com/go-redis/redis/v8"
)

type redisUserRepository struct {
	rclient  *redis.Client
	userRepo UserRepository
}

// NewRedisUserRepository ...
func NewRedisUserRepository(rclient *redis.Client, userRepo UserRepository) UserRepository {
	return &redisUserRepository{
		rclient:  rclient,
		userRepo: userRepo,
	}
}

// NewUser ...
func (r *redisUserRepository) NewUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	return r.userRepo.NewUser(ctx, user)
}

// GetUser ...
func (r *redisUserRepository) GetUser(ctx context.Context, uid string) (ruser *model.User, err error) {
	return r.userRepo.GetUser(ctx, uid)
}

// UpdateUser ...
func (r *redisUserRepository) UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	return r.userRepo.UpdateUser(ctx, user)
}

// DeleteUser ...
func (r *redisUserRepository) DeleteUser(ctx context.Context, uid string) (err error) {
	return r.userRepo.DeleteUser(ctx, uid)
}
