package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPopulation_GetSize(t *testing.T) {
	population := NewPopulation(make([]*Chromosome, 3))
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
	population := GenerateInitialPopulation(size, locationStartingPoint, locations)

	t.Run("should generate population with correct size", func(t *testing.T) {
		assert.Equal(t, size, population.GetSize())
	})
}

func TestPopulation_EvaluateFitness(t *testing.T) {
	defer t.Cleanup(cleanCache)
	calculatorMock := &mockDistanceCalculator{}
	calculatorMock.On("CalculateDistance", mock.Anything, mock.Anything).Return(2.0)
	locationStartingPoint := NewLocation(1, "any")
	locations := []*Location{
		NewLocation(2, "any2"),
		NewLocation(3, "any3"),
		NewLocation(4, "any4"),
		NewLocation(5, "any5"),
	}
	size := 4
	population := GenerateInitialPopulation(size, locationStartingPoint, locations)

	population.EvaluateFitness(calculatorMock)

	for _, chromosome := range population.Chromosomes {
		assert.NotEmpty(t, chromosome.Fitness)
	}
}
