package proxy

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/redis"
)

// NewRedisClient returns a new Redis client ready for use
func NewRedisClient(address string) (*redis.Client, error) {
	client, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error: could not connect to Redis server")
		return nil, err
	}
	return client, nil
}
