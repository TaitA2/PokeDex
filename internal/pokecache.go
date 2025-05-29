package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	entry map[string]cacheEntry
	mu *sync.Mutex
}

func NewCache(d time.Duration) Cache {
	var c Cache
	go c.reapLoop(d)
	return c

}

type cacheEntry struct{
	createdAt time.Time // time the entry was created
	val []byte // raw data
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{time.Now(), val}
}

func (c Cache) Get(key string) ([]byte , bool) {
	entry, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(d time.Duration) {
	time.Sleep(d)
	for k, entry := range c.entry {
		c.mu.Lock()
		defer c.mu.Unlock()
		if time.Since(entry.createdAt) > d {
			delete(c.entry, k)
		}
	}
}
