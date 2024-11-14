package core

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache_StoreAndRetrieve(t *testing.T) {
	cache := new(Cache)

	fromID := 1
	destinationID := 2
	expectedDistance := 22.765

	cache.CacheDistance(fromID, destinationID, expectedDistance)

	actualDistance, found := cache.GetFromCache(fromID, destinationID)

	assert.True(t, found)
	assert.Equal(t, expectedDistance, actualDistance)
}

func TestCache_EmptyCache(t *testing.T) {
	cache := new(Cache)

	fromID := 1
	destinationID := 2
	actualDistance, found := cache.GetFromCache(fromID, destinationID)
	assert.False(t, found)
	assert.Equal(t, 0.0, actualDistance)

	fromID = 999
	destinationID = 1000
	actualDistance, found = cache.GetFromCache(fromID, destinationID)
	assert.False(t, found)
	assert.Equal(t, 0.0, actualDistance)
}

func TestCache_Singleton(t *testing.T) {
	cache1 := GetCacheInstance()
	cache2 := GetCacheInstance()

	assert.Same(t, cache1, cache2)
}

func TestCache_OverwriteExistingValue(t *testing.T) {
	cache := new(Cache)

	fromID := 1
	destinationID := 2
	firstDistance := 10.5
	secondDistance := 20.75

	cache.CacheDistance(fromID, destinationID, firstDistance)

	actualDistance, found := cache.GetFromCache(fromID, destinationID)
	assert.True(t, found)
	assert.Equal(t, firstDistance, actualDistance)

	cache.CacheDistance(fromID, destinationID, secondDistance)

	actualDistance, found = cache.GetFromCache(fromID, destinationID)
	assert.True(t, found)
	assert.Equal(t, secondDistance, actualDistance)
}

func TestCache_Concurrency(t *testing.T) {
	cache := new(Cache)

	var wg sync.WaitGroup
	numRequests := 100

	for i := range numRequests {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fromID := i
			destinationID := i + 1
			distance := float64(i * 2)

			cache.CacheDistance(fromID, destinationID, distance)

			actualDistance, found := cache.GetFromCache(fromID, destinationID)
			assert.True(t, found)
			assert.Equal(t, distance, actualDistance)
		}(i)
	}

	wg.Wait()
}

func TestGenerateCacheKey(t *testing.T) {
	testCases := []struct {
		expected              string
		fromID, destinationID int
	}{
		{"2-9", 2, 9},
		{"4-1", 4, 1},
		{"238-981", 238, 981},
	}
	for _, testCase := range testCases {
		generatedKey := generateCacheKey(testCase.fromID, testCase.destinationID)
		assert.Equal(t, testCase.expected, generatedKey)
	}
}
