package dao

import (
	"context"
	"easygo/internal/entity"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// UserDaoSet 注入 DI
var UserDaoSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

// UserDao 表相关的数据库操作
type UserDao struct {
	DB *gorm.DB
}

type UserInsertParams struct {
	Username string
	Password string
	Personal string
}

type UserSelectParams struct {
	Username string
	Password string
}

// Insert 根据给定条件插入用户
func (u *UserDao) Insert(ctx context.Context, data UserInsertParams) error {

	db := u.DB.Model(&entity.User{})

	user := entity.User{
		Username: data.Username,
		Password: data.Password,
		Personal: data.Personal,
	}

	return db.Create(&user).Error
}

func (u *UserDao) Select(ctx context.Context, data UserSelectParams) (*entity.User, error) {

	db := u.DB.Model(&entity.User{})
	user := entity.User{}
	if data.Username != "" {
		db = db.Where("username=?", data.Username)
	}

	if data.Password != "" {
		db = db.Where("password=?", data.Password)
	}
	err := db.Find(&user).Error
	if err != nil {
		err = errors.Wrap(err, "Dao User Select")
	}

	return &user, err
}
