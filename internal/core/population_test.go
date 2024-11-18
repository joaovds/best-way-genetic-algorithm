package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPopulation_GetSize(t *testing.T) {
	population := NewPopulation(make([]*Chromosome, 3), MockGetCacheInstanceFn)
	assert.Equal(t, 3, population.GetSize())
}

func TestGenerateInitialPopulation(t *testing.T) {
	locationStartingPoint := NewLocation(1, "any")
	locations := []*Location{
		NewLocation(2, "any2"),
		NewLocation(3, "any3"),
		NewLocation(4, "any4"),
		NewLocation(5, "any5"),
	}

	size := 10
	population := GenerateInitialPopulation(size, locationStartingPoint, locations, MockGetCacheInstanceFn)

	t.Run("should generate population with correct size", func(t *testing.T) {
		assert.Equal(t, size, population.GetSize())
	})
}

func TestPopulation_EvaluateFitness(t *testing.T) {
	t.Run("calculates fitness for all chromosomes in the population", func(t *testing.T) {
		calculatorMock := &MockDistanceCalculator{}
		calculatorMock.On("CalculateDistance", mock.Anything, mock.Anything).Return(2.0)
		locationStartingPoint := NewLocation(1, "any")
		locations := []*Location{
			NewLocation(2, "any2"),
			NewLocation(3, "any3"),
			NewLocation(4, "any4"),
			NewLocation(5, "any5"),
		}
		size := 4
		population := GenerateInitialPopulation(size, locationStartingPoint, locations, MockGetCacheInstanceFn)

		population.EvaluateFitness(calculatorMock)

		var expectedTotalFitness float64
		t.Run("check if the fitness of all chromosomes is not empty", func(t *testing.T) {
			for _, chromosome := range population.Chromosomes {
				assert.NotEmpty(t, chromosome.Fitness)
				expectedTotalFitness += chromosome.Fitness
			}
		})

		t.Run("check if the total fitness sum is correct", func(t *testing.T) {
			assert.Equal(t, expectedTotalFitness, population.TotalFitness)
		})
	})
}

func TestPopulation_SortByFitness(t *testing.T) {
	chromosome1 := NewChromosome(NewGene(1, "any"), []*Gene{
		NewGene(2, "any2"),
	})
	chromosome1.Fitness = 7
	chromosome2 := NewChromosome(NewGene(1, "any"), []*Gene{
		NewGene(2, "any2"),
	})
	chromosome2.Fitness = 20
	chromosome3 := NewChromosome(NewGene(1, "any"), []*Gene{
		NewGene(2, "any2"),
	})
	chromosome3.Fitness = 12
	population := NewPopulation([]*Chromosome{
		chromosome1,
		chromosome2,
		chromosome3,
	}, MockGetCacheInstanceFn)

	population.SortByFitness()

	assert.Same(t, chromosome2, population.Chromosomes[0])
	assert.Same(t, chromosome3, population.Chromosomes[1])
	assert.Same(t, chromosome1, population.Chromosomes[2])
}
