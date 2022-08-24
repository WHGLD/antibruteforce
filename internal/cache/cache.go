package cache

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func (c *RedisCache) Get(ctx context.Context, key string) (int, error) {
	val, err := c.client.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return 0, errors.New("key does not exist")
	case err != nil:
		return 0, err
	case val == "":
		return 0, errors.New("value is empty")
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.New("value is not an int")
	}

	return intVal, nil
}

func (c *RedisCache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *RedisCache) Close() error {
	return c.client.Close()
}

func New(host string, port int, password string) (*RedisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisCache{client: rdb}, nil
}
