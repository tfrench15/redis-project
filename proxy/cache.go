package proxy

import (
	"log"
	"time"

	"github.com/hashicorp/golang-lru"
)

type Cache struct {
	cache  *lru.Cache
	expiry time.Duration
}

type CachedItem struct {
	value     string
	createdAt time.Time
}

func NewCache(size int, expiry time.Duration) *Cache {
	lru, err := lru.New(size)
	if err != nil {
		log.Fatal("Error creating cache")
	}
	return &Cache{
		cache:  lru,
		expiry: expiry,
	}
}

func (c *Cache) IsExpired(ci CachedItem) bool {
	if time.Now().Sub(ci.createdAt) < c.expiry {
		return false
	}
	return true
}
