package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGene(t *testing.T) {
	gene := NewGene(1, "any_address")
	assert.Equal(t, 1, gene.GetID())
	assert.Equal(t, "any_address", gene.Address)
	assert.Equal(t, 0.0, gene.Distance)
}

func TestCalculateDistanceToDestination(t *testing.T) {
	t.Run("should use cache if distance is already calculated", func(t *testing.T) {
		mockCache := MockGetCacheInstanceFn()
		for i := range 10 {
			from := NewGene(i, "any_address")
			destination := NewGene(i+1, "another_address")
			mockCache.CacheDistance(i, i+1, float64(i*99+2), 22+i*2)

			result, duration := from.CalculateDistanceToDestination(destination, mockCache)
			assert.Equal(t, float64(i*99+2), result)
			assert.Equal(t, 22+i*2, duration)
		}
	})
}
