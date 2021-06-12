package controller

import (
	"easygo/pkg/wrapper"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/wire"
)

// Register DI
var TestSet = wire.NewSet(wire.Struct(new(Test), "*"))

//
type Test struct {
	RedisTest *redis.Client
}

//jwt登录状态测试
func (t *Test) JwtTest(c *gin.Context) {

	username := c.GetString("username")

	wrapper.ResSuccess(c, username)
}

//redis存储测试
func (t *Test) RedisSetTest(c *gin.Context) {
	username := "xiao"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJePxk"
	err := t.RedisTest.Set(username, token, 0).Err()
	if err != nil {
		wrapper.ResError(c, wrapper.ErrRedisSet)
		return
	}
	wrapper.ResSuccess(c, nil)
}

//redis读取测试
func (t *Test) RedisGetTest(c *gin.Context) {
	val, err := t.RedisTest.Get("xiao").Result()
	if err != nil {
		wrapper.ResError(c, wrapper.ErrRedisGet)
		return
	}
	wrapper.ResSuccess(c, val)
}
