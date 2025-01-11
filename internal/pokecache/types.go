package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

}

func (c *Cache) Get(key string) (val []byte, wasFound bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if val, ok := c.cache[key]; ok {
		return val.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.cache {

			if time.Since(entry.createdAt) > interval {

				delete(c.cache, key)

			}

		}
		c.mu.Unlock()
	}
}
