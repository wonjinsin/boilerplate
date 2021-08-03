package repository

import (
	"fmt"

	"pikachu/config"
	"pikachu/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repository ...
type Repository struct {
	User UserRepository
}

// Init ...
func Init(pikachu *config.ViperConfig) (*Repository, error) {
	mysqlConn, err := mysqlConnect(pikachu)

	if err != nil {
		return nil, err
	}

	userRepo := NewGormUserRepository(mysqlConn)

	return &Repository{User: userRepo}, nil
}

func mysqlConnect(pikachu *config.ViperConfig) (mysql *gorm.DB, err error) {
	mysql, err = gorm.Open(getDialector(pikachu), &gorm.Config{})

	return mysql, err
}

func getDialector(pikachu *config.ViperConfig) gorm.Dialector {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=UTC",
		pikachu.GetString("database.username"),
		pikachu.GetString("database.password"),
		pikachu.GetString("database.host"),
		pikachu.GetInt("database.port"),
		pikachu.GetString("database.dbname"),
	)

	return mysql.Open(dbURI)
}

// UserRepository ...
type UserRepository interface {
	NewUser() *model.User
}
