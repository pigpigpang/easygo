package model

import (
	"context"
	"easygo/internal/dao"
	"easygo/internal/schema"
	"easygo/pkg/jwt"
	"github.com/google/wire"
)

var UserModelSet = wire.NewSet(wire.Struct(new(UserModel), "*"))

// UserModel 处理登录的主要逻辑
type UserModel struct {
	UserDao *dao.UserDao
	Jwt     *jwt.Jwt
}

func (u *UserModel) Register(ctx context.Context, data schema.RegisterReqBodySchema) (*schema.RegisterResBodySchema, error) {

	registerParams := dao.UserInsertParams{
		Username: data.Username,
		Password: data.Password,
		Personal: data.Personal,
	}

	err := u.UserDao.Insert(ctx, registerParams)

	if err != nil {
		return nil, err
	}

	registerRes := schema.RegisterResBodySchema{
		Msg: "注册成功",
	}

	return &registerRes, nil
}

func (u *UserModel) Login(ctx context.Context, data schema.LoginReqBodySchema) (*schema.LoginResBodySchema, error) {

	loginParams := dao.UserSelectParams{
		Username: data.Username,
		Password: data.Password,
	}

	user, err := u.UserDao.Select(ctx, loginParams)

	if err != nil {
		return nil, err
	}

	token, err := u.Jwt.Auth(data.Username)

	loginRes := schema.LoginResBodySchema{
		Msg:      "登录成功",
		Personal: user.Personal,
		Token:    token,
	}

	return &loginRes, nil
}
