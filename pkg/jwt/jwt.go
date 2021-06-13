package jwt

import (
	"easygo/internal/common/config"
	"easygo/pkg/wrapper"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"time"
)

var JWTSet = wire.NewSet(wire.Struct(new(Jwt), "*"))

// Jwt JWT
type Jwt struct {
	Redis *redis.Client
}

// Auth 发布签名
func (j *Jwt) Auth(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      config.C.JWT.Iss,
		"exp":      time.Now().Add(time.Second * time.Duration(config.C.JWT.ExpTime)).Unix(),
		"iat":      time.Now().Unix(),
		"username": username,
	})

	tokenString, err := token.SignedString([]byte(config.C.JWT.Secret))
	if err != nil {
		return "", wrapper.ErrJwtAuthToken
	}

	err = j.Redis.Set(username, tokenString, 0).Err()
	if err != nil {
		return "", wrapper.ErrRedisSet
	}

	return tokenString, nil
}

// Parse 验证签名
func (j *Jwt) Parse(tokenString string) (interface{}, error) {

	secret := config.C.JWT.Secret

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, wrapper.ErrInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		oldToken, err := j.Redis.Get(claims["username"].(string)).Result()
		if err != nil {
			return nil, err
		}
		if oldToken != tokenString {
			return nil, wrapper.ErrJwtBad
		}
		return claims["username"], nil
	}

	return nil, wrapper.ErrJwtExp
}
