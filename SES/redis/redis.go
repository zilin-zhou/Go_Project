package redis

import (
	"github.com/go-redis/redis/v8"
	"mail/config"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(
		&redis.Options{
			Addr:     config.SesSecret.GetString("Redis.Addr"),
			Password: config.SesSecret.GetString("Redis.Password"),
			DB:       0,
		})
}
