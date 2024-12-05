package redisrepo

import (
	"fmt"

	"github.com/calango-productions/api/internal/core/ports"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type RedisRepository struct {
	client *redis.Client
}

func New(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Store(conf ports.RedisStoreConf) error {
	dataStr, ok := conf.Data.(string)
	if !ok {
		return fmt.Errorf("invalid data type, expected string")
	}

	err := r.client.Set(context.Background(), conf.Key, dataStr, conf.Expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to store data in Redis: %v", err)
	}

	return nil
}

func (r *RedisRepository) Rescue(conf ports.RedisRescueConf) (any, error) {
	data, err := r.client.Get(context.Background(), conf.Key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("data not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get data from Redis: %v", err)
	}

	return data, nil
}
