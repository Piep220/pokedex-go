package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries 	map[string]cacheEntry
	mu 			*sync.RWMutex
	interval 	time.Duration //reap loop time
	ttl 		time.Duration //time to live, 0 for inf
	stopCh		chan struct{}
}

type cacheEntry struct {
	createdAt   time.Time
	lastUsed    time.Time
	value 		[]byte
}

//Creates new Cache. Interval is how often cleanup runs. If ttl=0, entries never expire.
func NewCache(interval, ttl time.Duration) Cache {
	c := Cache{
		entries:  make(map[string]cacheEntry),
		mu:		  &sync.RWMutex{},
		interval: interval,
		ttl: 	  ttl,
		stopCh:   make(chan struct{}),
	}

	go c.reapLoop()

	return c
}

//Stores a copy of val under key in cacheEntry.
func (c *Cache)Add(key string, val []byte){
	cp := make([]byte, len(val))
    copy(cp, val)

	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	c.entries[key] = cacheEntry{
		createdAt: now,
		lastUsed: now,
		value: cp,
	}
}

//Return copy of cached value and true if found. False if not found.
func (c *Cache)Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	e.lastUsed = time.Now()
	c.entries[key] = e

	out := make([]byte, len(e.value))
    copy(out, e.value)
	return out, true
}

//Background loop, removes stale data.
func (c *Cache)reapLoop(){
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
        case <-ticker.C:
            c.reap()
        case <-c.stopCh:
            return
        }
	}
}

//Removes expired data, older and last used than TTL.
func (c *Cache)reap() {
	if c.ttl == 0 {
		return
	}

	now := time.Now()

	c.mu.Lock()
	defer c.mu.Unlock()

	for k, e := range c.entries {
		if now.Sub(e.lastUsed) > c.ttl {
			delete(c.entries, k)
		}
		if now.Sub(e.createdAt) > c.ttl * 10 {
			delete(c.entries, k)
		}
	}
}

//Stop terminates background goroutine reapLoop.
func (c *Cache) Stop() {
    close(c.stopCh)
}
