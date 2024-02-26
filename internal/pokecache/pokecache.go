package pokecache

import (
	"sync"
	"time"
)
type Cache struct {
    data map[string]CacheEntry
    mu *sync.Mutex
}

type CacheEntry struct {
    createdAt time.Time
    Val []byte
}

var  Interval time.Duration 

func NewCache(dur time.Duration) Cache {
    c :=  Cache{
        data: make(map[string]CacheEntry),
		mu:   &sync.Mutex{},
    }
    go c.reapLoop(dur)
    Interval = dur
    return c
}

func(c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = CacheEntry{
        createdAt: time.Now().UTC(),
        Val: val,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.data[key]
	return val.Val, ok
}

func(c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(time.Second)
    for range ticker.C {
		c.reap(interval)	
	}
}

func(c *Cache) reap(interval time.Duration) {
            dt := time.Now().UTC().Add(-Interval)
            for k ,v := range c.data {
                if v.createdAt.Before(dt) {
                        c.mu.Lock()
                        delete(c.data, k)
                        c.mu.Unlock()
                }
            }
}
