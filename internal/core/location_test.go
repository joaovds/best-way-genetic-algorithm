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
	assert.IsType(t, &Gene{}, locationToGene)
}

func TestLocationsToGenes(t *testing.T) {
	locations := []*Location{
		NewLocation(1, "any1"),
		NewLocation(2, "any2"),
		NewLocation(3, "any3"),
	}
	genes := LocationsToGenes(locations)
	assert.Len(t, genes, len(locations))
	for i := range len(locations) {
		assert.Equal(t, locations[i].ID, genes[i].GetID())
		assert.Equal(t, locations[i].Address, genes[i].Address)
		assert.Zero(t, genes[i].Distance)
	}
}
