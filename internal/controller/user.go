package controller

import (
	"easygo/internal/model"
	"easygo/internal/schema"
	"easygo/pkg/logger"
	"easygo/pkg/wrapper"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User Login 登录结构体
type User struct {
	UserModel *model.UserModel
}

func (u *User) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.RegisterReqBodySchema
	if err := wrapper.ParseJSON(c, &data); err != nil {
		wrapper.ResError(c, err)
		return
	}

	res, err := u.UserModel.Register(ctx, data)
	if err != nil {
		wrapper.ResError(c, err)
		return
	}

	ctx = logger.NewTagContext(ctx, "__register__")
	logger.WithContext(ctx).Info("注册成功")

	wrapper.ResSuccess(c, res)
}

func (u *User) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.LoginReqBodySchema
	if err := wrapper.ParseJSON(c, &data); err != nil {
		wrapper.ResError(c, err)
		return
	}

	res, err := u.UserModel.Login(ctx, data)
	if err != nil {
		wrapper.ResError(c, err)
		return
	}

	ctx = logger.NewTagContext(ctx, "__login__")
	logger.WithContext(ctx).Info("登录成功")

	wrapper.ResSuccess(c, res)
}
