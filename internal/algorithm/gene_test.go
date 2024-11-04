package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGene(t *testing.T) {
	t.Run("CalculateDistance()", func(t *testing.T) {
		t.Run("should return 0.2 value", func(t *testing.T) {
			gene := NewGene(1, 0.1, 0.2)
			geneDestination := NewGene(3, 0.3, 0.4)

			distance := gene.CalculateDistanceToDestination(geneDestination)
			assert.Equal(t, 0.2, distance)
		})
	})
}
