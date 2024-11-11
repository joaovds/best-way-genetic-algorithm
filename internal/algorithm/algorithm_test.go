package algorithm

import (
	"testing"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestNewAlgorithm(t *testing.T) {
	t.Run("should return the population size as the factorial of the number less than or equal to 5", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			startingPoint := new(core.Location)
			locations := make([]*core.Location, 3)
			algorithm := NewAlgorithm(startingPoint, locations)
			assert.Equal(t, 6, algorithm.populationSize)
		})
	})

	startingPoint := new(core.Location)

	tests := []struct {
		name            string
		locations       []*core.Location
		expectedPopSize int
	}{
		{
			name:            "chromosomeSize less than or equal to 5",
			locations:       make([]*core.Location, 4),
			expectedPopSize: 24,
		},
		{
			name:            "chromosomeSize less than or equal to 5",
			locations:       make([]*core.Location, 5),
			expectedPopSize: 120,
		},
		{
			name:            "chromosomeSize greater than 5",
			locations:       make([]*core.Location, 60),
			expectedPopSize: 6000,
		},
		{
			name:            "populationSize exceeds MAX_POPULATION_SIZE",
			locations:       make([]*core.Location, 82),
			expectedPopSize: MAX_POPULATION_SIZE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			algorithm := NewAlgorithm(startingPoint, tt.locations)
			if algorithm.populationSize != tt.expectedPopSize {
				t.Errorf("expected populationSize %d, got %d", tt.expectedPopSize, algorithm.populationSize)
			}
		})
	}
}
