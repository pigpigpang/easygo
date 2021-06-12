package router

import (
	"easygo/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterAPI 路由列表
func (a *Router) registerAPI(app *gin.Engine) {
	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		v1.POST("/login", a.UserAPI.Login)
		v1.POST("/register", a.UserAPI.Register)
	}

	//测试组
	t1 := g.Group("/test")
	{
		t1.GET("/jwt", middleware.Auth(a.Jwt), a.TestAPI.JwtTest)
		t1.GET("/redis/set", a.TestAPI.RedisSetTest)
		t1.GET("/redis/get", a.TestAPI.RedisGetTest)
	}
}
