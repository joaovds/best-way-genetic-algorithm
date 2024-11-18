package distance

import (
	"testing"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestCalculateDistances(t *testing.T) {
	location1 := core.NewLocation(1, "any1")
	location2 := core.NewLocation(2, "any2")
	location3 := core.NewLocation(3, "any3")
	locations := []*core.Location{location1, location2, location3}

	cache := core.MockGetCacheInstanceFn()
	calculator := NewInBatchCalculator()

	calculator.CalculateDistances(locations, cache)
	checkDistanceInCache := func(id1, id2 int) {
		distance, exists := cache.GetFromCache(id1, id2)
		assert.True(t, exists)
		assert.NotZero(t, distance)
	}

	checkDistanceInCache(location1.ID, location2.ID)
	checkDistanceInCache(location1.ID, location3.ID)
	checkDistanceInCache(location2.ID, location1.ID)
	checkDistanceInCache(location2.ID, location3.ID)
	checkDistanceInCache(location3.ID, location1.ID)
	checkDistanceInCache(location3.ID, location2.ID)
}
