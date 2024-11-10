package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChromosome(t *testing.T) {
	startingPointGene := NewGene(1, "any_adress")
	genes := []Gene{
		NewGene(2, "any_adress2"),
		NewGene(3, "any_adress3"),
		NewGene(4, "any_adress4"),
	}

	chromosome := NewChromosome(startingPointGene, genes)
	assert.Equal(t, startingPointGene.GetID(), chromosome.StartingPoint.GetID())
	assert.Equal(t, startingPointGene.Address, chromosome.StartingPoint.Address)
	assert.Equal(t, startingPointGene.Distance, chromosome.StartingPoint.Distance)
	assert.Equal(t, startingPointGene.startingPoint, chromosome.StartingPoint.startingPoint)
	for i, gene := range genes {
		assert.Equal(t, gene.GetID(), chromosome.Genes[i].GetID())
		assert.Equal(t, gene.Address, chromosome.Genes[i].Address)
		assert.Equal(t, gene.Distance, chromosome.Genes[i].Distance)
		assert.Equal(t, gene.startingPoint, chromosome.Genes[i].startingPoint)
	}
}
