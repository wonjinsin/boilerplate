package repository

import (
	"context"
	"encoding/json"
	"pikachu/model"

	"github.com/go-redis/redis/v8"
)

type redisUserRepository struct {
	client   *redis.Client
	userRepo UserRepository
}

// NewRedisUserRepository ...
func NewRedisUserRepository(client *redis.Client, userRepo UserRepository) UserRepository {
	return &redisUserRepository{
		client:   client,
		userRepo: userRepo,
	}
}

// NewUser ...
func (r *redisUserRepository) NewUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	return r.userRepo.NewUser(ctx, user)
}

// GetUser ...
func (r *redisUserRepository) GetUser(ctx context.Context, uid string) (ruser *model.User, err error) {
	userJSON, err := r.client.Get(ctx, "dantats:user").Bytes()
	err = json.Unmarshal(userJSON, &ruser)
	if err == redis.Nil {
		return r.userRepo.GetUser(ctx, uid)
	} else if err != nil {
		return nil, err
	}

	return ruser, nil
}

// GetUserByEmail ...
func (r *redisUserRepository) GetUserByEmail(ctx context.Context, email string) (ruser *model.User, err error) {
	return r.userRepo.GetUserByEmail(ctx, email)
}

// UpdateUser ...
func (r *redisUserRepository) UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	return r.userRepo.UpdateUser(ctx, user)
}

// DeleteUser ...
func (r *redisUserRepository) DeleteUser(ctx context.Context, uid string) (err error) {
	return r.userRepo.DeleteUser(ctx, uid)
}
