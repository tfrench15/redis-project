package proxy

import (
	"log"
	"time"

	"github.com/hashicorp/golang-lru"
)

// Cache implements Hashicorp's LRU cache with an expiry field
// that is used to expire keys past a certain time.
type Cache struct {
	lru    *lru.Cache
	expiry time.Duration
}

// CachedItem is the object to be cached, where value is the
// value stored in Redis and createdAt is used to expire keys.
type CachedItem struct {
	value     string
	createdAt time.Time
}

// NewCache returns a new cache to work with
func NewCache(size int, expiry time.Duration) *Cache {
	lru, err := lru.New(size)
	if err != nil {
		log.Fatal("Error creating cache")
	}
	return &Cache{
		lru:    lru,
		expiry: expiry,
	}
}

// IsExpired checks whether a given item in the cache is stale
// or fresh.
func (c *Cache) IsExpired(ci CachedItem) bool {
	if time.Now().Sub(ci.createdAt) < c.expiry {
		return false
	}
	return true
}

// Lookup checks whether a key:value is cached and fresh.
// If the key is cached, Lookup returns the value and true,
// the empty string and false otherwise.
func (c *Cache) Lookup(key string) (string, bool) {
	v, ok := c.lru.Get(key)
	if !ok {
		return "", false
	}
	item := v.(CachedItem)
	if c.IsExpired(item) {
		return "", false
	}
	value := item.value
	return value, true
}
