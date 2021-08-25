package service

import (
	"log"
	"os"
	"pikachu/repository"
	"pikachu/util"

	"pikachu/model"
)

var zlog *util.Logger

func init() {
	_, err := util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[service] err[%s]", err.Error())
		os.Exit(1)
	}
}

// Service ...
type Service struct {
	User UserService
}

// Init ...
func Init(repo *repository.Repository) (*Service, error) {
	userSvc := NewUserService(repo.User)

	return &Service{User: userSvc}, nil
}

// UserService ...
type UserService interface {
	NewUser() *model.User
}
