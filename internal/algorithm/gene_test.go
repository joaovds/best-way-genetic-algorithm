package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGene(t *testing.T) {
	t.Run("CalculateDistance()", func(t *testing.T) {
		t.Run("should return a random number like distance", func(t *testing.T) {
			start := NewGene(1, 0.1, 0.2)
			destination := NewGene(3, 0.3, 0.4)
			distance := start.CalculateDistanceToDestination(destination)
			assert.NotNil(t, distance)
			assert.NotEmpty(t, distance)
		})

		t.Run("should return the same distance for the same genes as start and destination", func(t *testing.T) {
			start := NewGene(1, 0.1, 0.2)
			destination := NewGene(3, 0.3, 0.4)
			distance := start.CalculateDistanceToDestination(destination)
			distance2 := start.CalculateDistanceToDestination(destination)
			assert.Equal(t, distance, distance2)
			destination2 := NewGene(2, 0.2, 0.4)
			distance3 := start.CalculateDistanceToDestination(destination2)
			distance4 := start.CalculateDistanceToDestination(destination2)
			assert.Equal(t, distance3, distance4)
		})

		t.Run("should return 0 as the distance if the id of the gene is the same as that of the target gene", func(t *testing.T) {
			start := NewGene(1, 0.1, 0.2)
			destination := start
			distance := start.CalculateDistanceToDestination(destination)
			assert.Equal(t, 0.0, distance)
		})
	})
}
