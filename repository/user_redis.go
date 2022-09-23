package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/louistwiice/go/fripclose/entity"
)

func (c *UserClient) SaveTokenInRedis(key, value string) (string, error) {
	ctx := context.Background()

	return c.redis.Set(ctx, key, value, 5*time.Minute).Result()
}

func (c *UserClient) GetTokenInRedis(key string) (string, error) {
	ctx := context.Background()

	val, err := c.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", entity.ErrNotFoundInRedis
	}

	return val, err
}
