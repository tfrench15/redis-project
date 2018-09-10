package proxy

import (
	"github.com/mediocregopher/radix.v2/redis"
)

type RedisProxy struct {
	cache *Cache
	redis *redis.Client
}

func NewRedisProxy() *RedisProxy {
	return nil
}
