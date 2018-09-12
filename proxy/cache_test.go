package proxy

import (
	"testing"
	"time"
)

// SetupAndSeedCache creates and seeds a cache for testing.
func SetupAndSeedCache(cap int, exp time.Duration) *Cache {
	c := NewCache(cap, exp)
	c.lru.Add("hello",
		CachedItem{
			value:     "world",
			createdAt: time.Now(),
		},
	)
	c.lru.Add(
		"hi",
		CachedItem{
			value:     "there",
			createdAt: time.Now(),
		},
	)
	return c
}
func TestIsExpired(t *testing.T) {
	c := SetupAndSeedCache(3, 3*time.Second)

	time.Sleep(1 * time.Second)

	res1, ok := c.lru.Get("hello")
	if !ok {
		t.Error("Error: could not find key")
	}
	item1 := res1.(CachedItem)
	exp1 := c.IsExpired(item1)
	if exp1 {
		t.Error("Error: fresh key is expired")
	}
	time.Sleep(4 * time.Second)

	res2, ok := c.lru.Get("hi")
	if !ok {
		t.Error("Error: could not find key")
	}
	item2 := res2.(CachedItem)
	exp2 := c.IsExpired(item2)
	if !exp2 {
		t.Error("Error: stale key not expired")
	}
}

func TestLookup(t *testing.T) {
	c := SetupAndSeedCache(5, 5*time.Second)

	tests := []struct {
		key    string
		value  string
		exists bool
	}{
		{"hello", "world", true},
		{"hi", "there", true},
		{"california", "", false},
	}

	for _, test := range tests {
		v, ok := c.Lookup(test.key)
		if (v != test.value) || (ok != test.exists) {
			t.Errorf("Error: expected value %v and bool %v, got value %v and bool %v", test.value, test.exists, v, ok)
		}
	}
}
