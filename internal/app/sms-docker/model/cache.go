package model

import (
	"fmt"

	"github.com/go-redis/redis"
)

type CacheConfig struct {
	Hostname string `json:"Hostname"` // Cache Server name
	Port     int    `json:"Port"`     // Cache port
}

type Cache interface {
	Increment(key string) (result int64, err error)
}

type RedisCache struct {
	Client *redis.Client
}

func InitRedisCache(c CacheConfig) (cache *RedisCache) {
	return &RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr: cacheAddress(c), // "redis:6379" when in container
		}),
	}
}

func (rc *RedisCache) Increment(key string) (result int64, err error) {
	return rc.Client.Incr(key).Result()
}

// cahceAddress returns the cache address
func cacheAddress(c CacheConfig) string {
	return c.Hostname + ":" + fmt.Sprintf("%d", c.Port)
}
