package proxy

import (
	"github.com/mediocregopher/radix.v2/redis"
)

// RedisProxy implements a cache and connection to the backing Redis
type RedisProxy struct {
	cache *Cache
	redis *redis.Client
}

// NewRedisProxy returns an instance of RedisProxy
func NewRedisProxy(c *Cache, rc *redis.Client) *RedisProxy {
	return &RedisProxy{
		cache: c,
		redis: rc,
	}
}
