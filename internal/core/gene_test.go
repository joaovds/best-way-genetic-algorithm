package core

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGene(t *testing.T) {
	gene := NewGene(1, "any_address")
	assert.Equal(t, 1, gene.GetID())
	assert.Equal(t, "any_address", gene.Address)
	assert.Equal(t, 0.0, gene.Distance)
	assert.False(t, gene.IsStartingPoint())

	t.Run("should return the starting point as true", func(t *testing.T) {
		gene.SetStartingPoint()
		assert.True(t, gene.IsStartingPoint())
	})
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

type mockDistanceCalculator struct {
	mock.Mock
}

func (m *mockDistanceCalculator) CalculateDistance(from, to *Gene) float64 {
	args := m.Called(from, to)
	return args.Get(0).(float64)
}

func TestCalculateDistanceToDestination(t *testing.T) {
	cleanCache := func() { distancesCache = make(map[string]float64) }

	t.Run("should return 0 if destinationID is equal to fromID", func(t *testing.T) {
		from := NewGene(1, "any_address")
		destination := from
		calculatorMock := &mockDistanceCalculator{}
		result := from.CalculateDistanceToDestination(&destination, calculatorMock)
		assert.Equal(t, 0.0, result)
		calculatorMock.AssertNotCalled(t, "CalculateDistance")
	})

	t.Run("should calculate distance if destinationID is different from fromID", func(t *testing.T) {
		from := NewGene(1, "any_address")
		destination := NewGene(2, "another_address")
		calculatorMock := &mockDistanceCalculator{}
		calculatorMock.On("CalculateDistance", &from, &destination).Return(17.77)

		result := from.CalculateDistanceToDestination(&destination, calculatorMock)
		assert.Equal(t, 17.77, result)
		calculatorMock.AssertCalled(t, "CalculateDistance", &from, &destination)
	})

	t.Run("should use cache if distance is already calculated", func(t *testing.T) {
		defer t.Cleanup(cleanCache)
		from := NewGene(1, "any_address")
		destination := NewGene(2, "another_address")
		calculatorMock := &mockDistanceCalculator{}

		cacheKey := generateCacheKey(from.id, destination.id)
		distancesCache[cacheKey] = 83.6

		result := from.CalculateDistanceToDestination(&destination, calculatorMock)
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
		calculatorMock.On("CalculateDistance", &from, &destination).Return(17.77)

		var counter int32
		var wg sync.WaitGroup
		numGoroutines := 10

		testFunc := func() {
			defer wg.Done()
			from.CalculateDistanceToDestination(&destination, calculatorMock)
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
