package repository

import (
	"context"
	"pikachu/model"
	"pikachu/util"

	"github.com/juju/errors"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	Conn *gorm.DB
}

// NewGormUserRepository ...
func NewGormUserRepository(conn *gorm.DB) UserRepository {
	migrations := []interface{}{
		&model.User{},
	}
	if err := conn.Set("gorm:table_options", util.DBCharsetOption).Migrator().AutoMigrate(migrations...); err != nil {
		zlog.Panicw("NewGormDealRepository Unable to AutoMigrate DealRepository", "err", err)
	}

	return &gormUserRepository{Conn: conn}
}

// NewUser ...
func (g *gormUserRepository) NewUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	zlog.With(ctx).Infow("Repository NewUser", "user", user)

	scope := g.Conn.WithContext(ctx)
	if err = scope.Create(&user).Error; err != nil {
		zlog.With(ctx).Errorw("Create User", "err", err)
		return nil, err
	}

	return user, nil
}

// GetUser ...
func (g *gormUserRepository) GetUser(ctx context.Context, uid string) (ruser *model.User, err error) {
	zlog.With(ctx).Infow("[New Repository Service]", "uid", uid)

	scope := g.Conn.WithContext(ctx)
	scope = scope.Where("users.uid = ?", uid).Find(&ruser)
	if err = scope.Error; err != nil {
		zlog.With(ctx).Errorw("Find User", "uid", uid, "err", err)
		return nil, err
	}
	if scope.RowsAffected == 0 {
		return nil, errors.UserNotFoundf("User is not exist")
	}

	return ruser, nil
}

// UpdateUser ...
func (g *gormUserRepository) UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	zlog.With(ctx).Infow("[New Repository Service]", "user", user)

	scope := g.Conn.WithContext(ctx)
	if err = scope.Updates(user).Error; err != nil {
		zlog.With(ctx).Errorw("Update User Failed", "user", user, "err", err)
		return nil, err
	}

	return user, nil
}

// DeleteUser ...
func (g *gormUserRepository) DeleteUser(ctx context.Context, uid string) (err error) {
	zlog.With(ctx).Infow("[New Repository Service]", "uid", uid)

	scope := g.Conn.WithContext(ctx)
	if err = scope.Where("uid = ?", uid).Delete(&model.User{}).Error; err != nil {
		zlog.With(ctx).Errorw("Delete User Failed", "uid", uid, "err", err)
		return err
	}

	return nil
}
