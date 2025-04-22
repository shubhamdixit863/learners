package internal

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func (r *RedisCache) SetData(key, value interface{}) {
	err := r.client.Set(context.Background(), key.(string), value, 0).Err()
	if err != nil {
		panic(err)
	}

}

func (r *RedisCache) GetData(key interface{}) (interface{}, error) {
	result, err := r.client.Get(context.Background(), key.(string)).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewRedisCache(client *redis.Client) Cache {
	return &RedisCache{
		client: client,
	}
}
