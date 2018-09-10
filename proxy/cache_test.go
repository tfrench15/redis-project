package proxy

import (
	"testing"
	"time"
)

func TestIsExpired(t *testing.T) {
	c := NewCache(3, 3*time.Second)

	c.cache.Add(
		"hello",
		CachedItem{
			value:     "world",
			createdAt: time.Now(),
		},
	)
	c.cache.Add(
		"hi",
		CachedItem{
			value:     "there",
			createdAt: time.Now(),
		},
	)
	time.Sleep(1 * time.Second)

	res1, ok := c.cache.Get("hello")
	if !ok {
		t.Error("Error: could not find key")
	}
	item1 := res1.(CachedItem)
	exp1 := c.IsExpired(item1)
	if exp1 {
		t.Error("Error: fresh key is expired")
	}
	time.Sleep(4 * time.Second)

	res2, ok := c.Cache.Get("hi")
	if !ok {
		t.Error("Error: could not find key")
	}
	item2 := res2.(CachedItem)
	exp2 := c.IsExpired(item2)
	if !exp2 {
		t.Error("Error: stale key not expired")
	}
}
