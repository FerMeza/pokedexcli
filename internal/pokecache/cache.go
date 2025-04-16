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
	content  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		content:  map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}
	cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.content[key] = cacheEntry{
		time.Now(),
		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	result, ok := c.content[key]
	if !ok {
		return []byte{}, false
	}
	return result.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for t := range ticker.C {
			c.mu.Lock()
			for key, value := range c.content {
				diff := t.Sub(value.createdAt)
				if diff.Seconds() >= c.interval.Seconds() {
					delete(c.content, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
