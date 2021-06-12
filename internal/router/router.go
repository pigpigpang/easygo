package router

import (
	"easygo/internal/controller"
	"easygo/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// RouterSet 路由注入
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

// Router 路由管理器
type Router struct {
	UserAPI *controller.User
	TestAPI *controller.Test
	Jwt     *jwt.Jwt
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) {
	a.registerAPI(app)
}
