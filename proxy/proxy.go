package proxy

import (
	"fmt"
	"time"

	"github.com/mediocregopher/radix.v2/redis"
)

// RedisProxy implements a cache and connection to the backing Redis
type RedisProxy struct {
	cache       *Cache
	redisClient *redis.Client
}

// NewRedisProxy returns an instance of RedisProxy
func NewRedisProxy(c *Cache, rc *redis.Client) *RedisProxy {
	return &RedisProxy{
		cache:       c,
		redisClient: rc,
	}
}

// FetchFromRedis looks up the given key in Redis.  It returns
// the value and true if found, the empty string and false otherwise.
func (rp *RedisProxy) FetchFromRedis(key string) (string, bool) {
	v, err := rp.redisClient.Cmd("GET", key).Str()
	if err != nil {
		fmt.Printf("Error fetching from Redis: %v", err)
		return "", false
	}
	if v == "" {
		fmt.Println("Error: key not found")
		return "", false
	}
	rp.cache.lru.Add(key, CachedItem{value: v, createdAt: time.Now()})
	return v, true
}
