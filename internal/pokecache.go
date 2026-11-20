package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	data map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        data: make(map[string]cacheEntry),
    }

    go c.reapLoop(interval)

    return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	n, ok := c.data[key]
	c.mu.Unlock()
	return n.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.mu.Lock()
		for key, entry := range c.data{
			if time.Since(entry.createdAt) > interval {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
