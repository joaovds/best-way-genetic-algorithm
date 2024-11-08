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
	assert.False(t, gene.IsStartingPoint())

	t.Run("should return the starting point as true", func(t *testing.T) {
		gene.SetStartingPoint()
		assert.True(t, gene.IsStartingPoint())
	})
}
