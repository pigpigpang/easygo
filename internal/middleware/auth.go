package middleware

import (
	"easygo/pkg/jwt"
	"easygo/pkg/wrapper"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(Jwt *jwt.Jwt) gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			wrapper.ResError(c, wrapper.ErrJwtEmpty)
			c.Abort()
			return
		}

		jwtStringArray := strings.Split(tokenString, " ")
		jwtStringArrayLen := len(jwtStringArray)
		if jwtStringArrayLen == 1 {
			wrapper.ResError(c, wrapper.ErrJwtEmpty)
			c.Abort()
			return
		}

		username, err := Jwt.Parse(jwtStringArray[1])
		if err != nil {
			wrapper.ResError(c, wrapper.ErrInvalidToken)
			c.Abort()
			return
		}

		newToken, err := Jwt.Auth(username.(string))
		if err != nil {
			wrapper.ResError(c, err)
			c.Abort()
			return
		}
		c.Header("rtr-token", newToken)

		c.Set("username", username.(string))
		c.Next()
	}
}
