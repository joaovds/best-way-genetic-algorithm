package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_ToNewGene(t *testing.T) {
	location := NewLocation(1, "any")
	locationToGene := location.ToNewGene()
	assert.Equal(t, location.ID, locationToGene.GetID())
	assert.Equal(t, location.Address, locationToGene.Address)
	assert.IsType(t, Gene{}, locationToGene)
}
