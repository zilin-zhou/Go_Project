package redis

import (
	"github.com/go-redis/redis/v8"
	"message/config"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(
		&redis.Options{
			Addr:     config.SmsSecret.GetString("Redis.Addr"),
			Password: config.SmsSecret.GetString("Redis.Password"),
			DB:       0,
		})
}
