package repository

import (
	"pikachu/model"

	"gorm.io/gorm"
)

type gormUserRepository struct {
	Conn *gorm.DB
}

// NewGormUserRepository ...
func NewGormUserRepository(conn *gorm.DB) UserRepository {
	return &gormUserRepository{Conn: conn}
}

// NewUser ...
func (g *gormUserRepository) NewUser() *model.User {
	return &model.User{}
}
