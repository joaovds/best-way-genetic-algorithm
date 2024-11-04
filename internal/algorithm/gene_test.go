package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGene(t *testing.T) {
	t.Run("CalculateDistance()", func(t *testing.T) {
		t.Run("should return a random number like distance", func(t *testing.T) {
			gene := NewGene(1, 0.1, 0.2)
			geneDestination := NewGene(3, 0.3, 0.4)

			distance := gene.CalculateDistanceToDestination(geneDestination)
			assert.NotNil(t, distance)
			assert.NotEmpty(t, distance)
		})

		t.Run("should return 0 as the distance if the id of the gene is the same as that of the target gene", func(t *testing.T) {
			gene := NewGene(1, 0.1, 0.2)
			geneDestination := gene

			distance := gene.CalculateDistanceToDestination(geneDestination)
			assert.Equal(t, 0.0, distance)
		})
	})
}
