package core

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGene(t *testing.T) {
	gene := NewGene(1, "any_address")
	assert.Equal(t, 1, gene.GetID())
	assert.Equal(t, "any_address", gene.Address)
	assert.Equal(t, 0.0, gene.Distance)
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

func TestCalculateDistanceToDestination(t *testing.T) {
	t.Run("should return 0 if destinationID is equal to fromID", func(t *testing.T) {
		defer t.Cleanup(cleanCache)
		from := NewGene(1, "any_address")
		destination := from
		calculatorMock := NewMockDistanceCalculator()
		result := from.CalculateDistanceToDestination(destination, calculatorMock)
		assert.Equal(t, 0.0, result)
		calculatorMock.AssertNotCalled(t, "CalculateDistance")
	})

	t.Run("should calculate distance if destinationID is different from fromID", func(t *testing.T) {
		defer t.Cleanup(cleanCache)
		from := NewGene(1, "any_address")
		destination := NewGene(2, "another_address")
		calculatorMock := NewMockDistanceCalculator()
		calculatorMock.On("CalculateDistance", from, destination).Return(17.77)

		result := from.CalculateDistanceToDestination(destination, calculatorMock)
		assert.Equal(t, 17.77, result)
		calculatorMock.AssertCalled(t, "CalculateDistance", from, destination)
	})

	t.Run("should use cache if distance is already calculated", func(t *testing.T) {
		defer t.Cleanup(cleanCache)
		from := NewGene(1, "any_address")
		destination := NewGene(2, "another_address")
		calculatorMock := &mockDistanceCalculator{}

		cacheKey := generateCacheKey(from.id, destination.id)
		distancesCache[cacheKey] = 83.6

		result := from.CalculateDistanceToDestination(destination, calculatorMock)
		assert.Equal(t, 83.6, result)
		calculatorMock.AssertNotCalled(t, "CalculateDistance")

		t.Cleanup(func() {
			distancesCache = make(map[string]float64)
		})
	})

	t.Run("should lock the cache and prevent concurrent access", func(t *testing.T) {
		defer t.Cleanup(cleanCache)
		from := NewGene(1, "any_address")
		destination := NewGene(2, "another_address")
		calculatorMock := &mockDistanceCalculator{}
		calculatorMock.On("CalculateDistance", from, destination).Return(17.77)

		var counter int32
		var wg sync.WaitGroup
		numGoroutines := 10

		testFunc := func() {
			defer wg.Done()
			from.CalculateDistanceToDestination(destination, calculatorMock)
			atomic.AddInt32(&counter, 1)
		}

		wg.Add(numGoroutines)
		for range numGoroutines {
			go testFunc()
		}
		wg.Wait()
		calculatorMock.AssertNumberOfCalls(t, "CalculateDistance", 1)
		assert.Equal(t, int32(numGoroutines), counter)
	})
}
