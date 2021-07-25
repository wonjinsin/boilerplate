package repository

import (
	"fmt"

	"pikachu/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repository ...
type Repository struct {
	User UserRepository
}

// Init ...
func Init() (*Repository, error) {
	mysqlConn, err := mysqlConnect()

	if err != nil {
		return nil, err
	}

	userRepo := NewGormUserRepository(mysqlConn)

	return &Repository{User: userRepo}, nil
}

// UserRepository ...
type UserRepository interface {
	NewUser() *model.User
}

func mysqlConnect() (mysql *gorm.DB, err error) {
	mysql, err = gorm.Open(getDialector(), &gorm.Config{})

	return mysql, err
}

func getDialector() gorm.Dialector {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=UTC",
		"root",
		"mysqlvotmdnjem",
		"127.0.0.1",
		53306,
		"pikachu",
	)

	return mysql.Open(dbURI)
}
