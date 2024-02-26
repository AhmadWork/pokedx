package pokecache

import (
	"fmt"
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
    fmt.Println(len(c.data))
    c.data[key] = CacheEntry{
        createdAt: time.Now().UTC(),
        Val: val,
    }
    fmt.Println(len(c.data))
}

func (c *Cache) Get(key string) ([]byte, bool) {
    for key, _ := range c.data {
        fmt.Println(key)
    }
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
