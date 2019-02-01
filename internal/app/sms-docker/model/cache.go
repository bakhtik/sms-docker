package model

import "github.com/go-redis/redis"

type Cache interface {
	Increment(key string) (result int64, err error)
}

type RedisCache struct {
	Client *redis.Client
}

func (rc *RedisCache) Increment(key string) (result int64, err error) {
	return rc.Client.Incr(key).Result()
}
