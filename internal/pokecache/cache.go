package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents a single entry in the cache
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache stores response data with timestamps
type Cache struct {
	entries map[string]cacheEntry
	mutex   sync.Mutex
}

// Add adds a new entry to the cache
func (c *Cache) Add(key string, val []byte) {
	// Lock the mutex before accessing the map
	c.mutex.Lock()
	// Make sure to unlock when we're done
	defer c.mutex.Unlock()

	// Now it's safe to modify the map
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves an entry from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	// Lock the mutex before accessing the map
	c.mutex.Lock()
	// Make sure to unlock when we're done
	defer c.mutex.Unlock()

	// Now it's safe to read from the map
	entry, found := c.entries[key]
	if !found {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			// Check if the entry is older than the interval
			if now.Sub(entry.createdAt) > interval {
				// Remove expired entries
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}

// NewCache creates a new cache with the given interval for cache entry expiration
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}

	// Start the background reaping process
	go cache.reapLoop(interval)

	return cache
}
