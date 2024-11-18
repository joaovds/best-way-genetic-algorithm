package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	cache := MockGetCacheInstanceFn()

	startingPointGene := NewGene(1, "any_adress")
	genes := []*Gene{
		NewGene(2, "any_adress2"),
		NewGene(3, "any_adress3"),
	}

	cache.CacheDistance(1, 2, 1.0)
	cache.CacheDistance(1, 3, 1.0)
	cache.CacheDistance(2, 1, 1.0)
	cache.CacheDistance(2, 3, 1.0)
	cache.CacheDistance(3, 1, 1.0)
	cache.CacheDistance(3, 2, 1.0)

	chromosome := NewChromosome(startingPointGene, genes)
	fitness := chromosome.CalculateFitness(cache)

	expectedFitness := float64(1) / float64(3)
	assert.Equal(t, chromosome.Fitness, fitness)
	assert.Equal(t, expectedFitness, fitness)
	for _, gene := range chromosome.Genes {
		assert.Equal(t, 1.0, gene.Distance)
	}
}
