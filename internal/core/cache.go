package core

import (
	"fmt"
	"sync"
)

var (
	cacheInstance *Cache
	once          sync.Once
)

type (
	Cache struct {
		cacheMap sync.Map
	}

	cacheEntry struct {
		distance float64
	}

	GetCacheInstanceFn func() *Cache
)

func (c *Cache) CacheDistance(fromID, destinationID int, distance float64) {
	cacheKey := generateCacheKey(fromID, destinationID)
	c.cacheMap.Store(cacheKey, cacheEntry{
		distance: distance,
	})
}

func (c *Cache) GetFromCache(fromID, destinationID int) (float64, bool) {
	cacheKey := generateCacheKey(fromID, destinationID)
	value, ok := c.cacheMap.Load(cacheKey)
	if !ok {
		return 0, false
	}

	entry := value.(cacheEntry)
	return entry.distance, true
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
