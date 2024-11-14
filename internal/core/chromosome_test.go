package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewChromosome(t *testing.T) {
	startingPointGene := NewGene(1, "any_adress")
	genes := []*Gene{
		NewGene(2, "any_adress2"),
		NewGene(3, "any_adress3"),
		NewGene(4, "any_adress4"),
	}

	chromosome := NewChromosome(startingPointGene, genes)
	assert.Equal(t, startingPointGene.GetID(), chromosome.StartingPoint.GetID())
	assert.Equal(t, startingPointGene.Address, chromosome.StartingPoint.Address)
	assert.Equal(t, startingPointGene.Distance, chromosome.StartingPoint.Distance)
	for i, gene := range genes {
		assert.Equal(t, gene.GetID(), chromosome.Genes[i].GetID())
		assert.Equal(t, gene.Address, chromosome.Genes[i].Address)
		assert.Equal(t, gene.Distance, chromosome.Genes[i].Distance)
	}
}

func TestChromosome_ShufflingGenes(t *testing.T) {
	genes := []*Gene{
		NewGene(1, "any_adress1"),
		NewGene(2, "any_adress2"),
		NewGene(3, "any_adress3"),
		NewGene(4, "any_adress4"),
	}
	chromosome := &Chromosome{Genes: genes}

	t.Run("should shuffle genes", func(t *testing.T) {
		originalOrder := make([]*Gene, len(chromosome.Genes))
		copy(originalOrder, chromosome.Genes)

		chromosome.ShufflingGenes()
		assert.ElementsMatch(t, originalOrder, chromosome.Genes, "The shuffled genes should contain the same elements as the original genes")
	})
}

func TestChromosome_CalculateFitness(t *testing.T) {
	defer t.Cleanup(cleanCache)
	calculatorMock := &mockDistanceCalculator{}
	calculatorMock.On("CalculateDistance", mock.Anything, mock.Anything).Return(2.0)

	startingPointGene := NewGene(1, "any_adress")
	genes := []*Gene{
		NewGene(2, "any_adress2"),
		NewGene(3, "any_adress3"),
		NewGene(4, "any_adress4"),
	}

	chromosome := NewChromosome(startingPointGene, genes)
	fitness := chromosome.CalculateFitness(calculatorMock)

	expectedFitness := 8.0
	assert.Equal(t, chromosome.Fitness, fitness)
	assert.Equal(t, expectedFitness, fitness)
	for _, gene := range chromosome.Genes {
		assert.Equal(t, 2.0, gene.Distance)
	}
}
