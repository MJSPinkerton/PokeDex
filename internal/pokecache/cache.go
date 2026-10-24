package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mtx   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	newCache := Cache{
		cache: map[string]cacheEntry{},
		mtx:   &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	entry, exists := c.cache[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.Tick(interval)
	for range ticker {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for key, val := range c.cache {
		if time.Since(val.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}
