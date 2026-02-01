package pokecache

import (
	"log/slog"
	"sync"
	"time"
)

type Cache struct {
	mu       *sync.Mutex
	interval time.Duration
	entries  map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	var mutex sync.Mutex
	newCache := Cache{
		entries:  map[string]cacheEntry{},
		interval: interval,
		mu:       &mutex,
	}

	go func() {
		for {
			_ = <-ticker.C
			newCache.reapLoop()
		}
	}()

	slog.Info("New cache created")
	return &newCache
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	slog.Info("Cache entry added")
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, exists := c.entries[key]
	if !exists {
		slog.Info("Cache miss")
		return nil, false
	}
	return val.val, true
}

func (c Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	slog.Info("Reap loop ran")
	for i, entry := range c.entries {
		if time.Since(entry.createdAt) > c.interval {
			delete(c.entries, i)
			slog.Info("Reaped", "url", i)
		}
	}
}
