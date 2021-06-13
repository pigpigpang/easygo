package internal

import (
	"easygo/internal/common/config"
	"github.com/go-redis/redis"
	"strconv"
)

// initRedis 初始化redis
func InitRedis() (*redis.Client, error) {
	c := config.C.REDIS
	if c.Enable == "false" {
		return nil, nil
	}
	client, err := NewRedis()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewRedis
func NewRedis() (*redis.Client, error) {

	Address := "localhost:" + strconv.Itoa(config.C.REDIS.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     Address,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
