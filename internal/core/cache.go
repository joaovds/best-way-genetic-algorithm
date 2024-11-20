package core

import (
	"fmt"
	"sync"
)

var (
	cacheInstance *Cache
	once          sync.Once
	cacheLocks    sync.Map
)

type (
	Cache struct {
		cacheMap sync.Map
	}

	cacheEntry struct {
		distance float64
		duration int
	}

	GetCacheInstanceFn func() *Cache
)

func (c *Cache) CacheDistance(fromID, destinationID int, distance float64, duration int) {
	cacheKey := generateCacheKey(fromID, destinationID)
	c.cacheMap.Store(cacheKey, cacheEntry{
		distance: distance,
		duration: duration,
	})
}

func (c *Cache) GetFromCache(fromID, destinationID int) (distance float64, duration int, ok bool) {
	cacheKey := generateCacheKey(fromID, destinationID)
	value, ok := c.cacheMap.Load(cacheKey)
	if !ok {
		return 0, 0, false
	}

	entry := value.(cacheEntry)
	return entry.distance, entry.duration, true
}

func (c *Cache) Clean() {
	c = new(Cache)
}

func generateCacheKey(fromID, destinationID int) string {
	return fmt.Sprintf("%d-%d", fromID, destinationID)
}

func GetCacheInstance() *Cache {
	once.Do(func() {
		cacheInstance = new(Cache)
	})
	return cacheInstance
}
