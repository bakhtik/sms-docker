package main

import "github.com/go-redis/redis"

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}
