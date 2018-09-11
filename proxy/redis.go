package proxy

import (
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

// NewRedisClient returns a new Redis client ready for use
func NewRedisClient(address string) *redis.Client {
	client, err := redis.Dial("tcp", address)
	if err != nil {
		log.Fatal("Error: could not connect to Redis server")
		return nil
	}
	return client
}
