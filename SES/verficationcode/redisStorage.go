package verficationcode

import (
	"context"
	v8 "github.com/go-redis/redis/v8"
	"mail/redis"
	"time"
)

type redisStorage struct {
	client *v8.Client
}

func NewRedisStorage() storge {

	return &redisStorage{
		client: redis.Client,
	}
}

func (r *redisStorage) Set(key, val string, duration time.Duration) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, val, duration).Err()
}
func (r *redisStorage) Get(key string) (string, error) {
	ctx := context.Background()
	return r.client.Get(ctx, key).Result()
}
