package middleware

import (
	"easygo/pkg/wrapper"
	"github.com/gin-gonic/gin"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		wrapper.ResError(c, wrapper.ErrMethodNotAllow)
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		wrapper.ResError(c, wrapper.ErrNotFound)
	}
}
