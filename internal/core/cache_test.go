package core

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache_StoreAndRetrieve(t *testing.T) {
	cache := new(Cache)

	fromID := 1
	destinationID := 2
	expectedDistance := 22.765

	cache.CacheDistance(fromID, destinationID, expectedDistance, 17)

	actualDistance, duration, found := cache.GetFromCache(fromID, destinationID)

	assert.True(t, found)
	assert.Equal(t, expectedDistance, actualDistance)
	assert.Equal(t, 17, duration)
}

func TestCache_EmptyCache(t *testing.T) {
	cache := new(Cache)

	fromID := 1
	destinationID := 2
	actualDistance, duration, found := cache.GetFromCache(fromID, destinationID)
	assert.False(t, found)
	assert.Equal(t, 0.0, actualDistance)
	assert.Equal(t, 0, duration)

	fromID = 999
	destinationID = 1000
	actualDistance, duration, found = cache.GetFromCache(fromID, destinationID)
	assert.False(t, found)
	assert.Equal(t, 0.0, actualDistance)
	assert.Equal(t, 0, duration)
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

	cache.CacheDistance(fromID, destinationID, firstDistance, 17)

	actualDistance, duration, found := cache.GetFromCache(fromID, destinationID)
	assert.True(t, found)
	assert.Equal(t, firstDistance, actualDistance)
	assert.Equal(t, 17, duration)

	cache.CacheDistance(fromID, destinationID, secondDistance, 22)

	actualDistance, duration, found = cache.GetFromCache(fromID, destinationID)
	assert.True(t, found)
	assert.Equal(t, secondDistance, actualDistance)
	assert.Equal(t, 22, duration)
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

			cache.CacheDistance(fromID, destinationID, distance, 22)

			actualDistance, duration, found := cache.GetFromCache(fromID, destinationID)
			assert.True(t, found)
			assert.Equal(t, distance, actualDistance)
			assert.Equal(t, 22, duration)
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

func TestCache(t *testing.T) {
	t.Run("should store and retrieve distance from cache", func(t *testing.T) {
		cache := MockGetCacheInstanceFn()
		from := NewGene(1, "address1")
		destination := NewGene(2, "address2")

		cache.CacheDistance(1, 2, 10.0, 22)
		distance, duration := from.CalculateDistanceToDestination(destination, cache)
		assert.Equal(t, 10.0, distance)
		assert.Equal(t, 22, duration)

		storedDistance, storedDuration, found := cache.GetFromCache(from.GetID(), destination.GetID())
		assert.True(t, found)
		assert.Equal(t, 10.0, storedDistance)
		assert.Equal(t, 22, storedDuration)
	})

	t.Run("should handle concurrent access correctly", func(t *testing.T) {
		cache := MockGetCacheInstanceFn()
		cache.CacheDistance(1, 2, 10, 22)

		var wg sync.WaitGroup
		numGoroutines := 10
		var counter int32

		testFunc := func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}

		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go testFunc()
		}

		wg.Wait()

		assert.Equal(t, int32(numGoroutines), counter)
	})
}
