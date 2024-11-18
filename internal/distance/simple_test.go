package distance

import (
	"testing"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestSimpleDistanceCalculator_CalculateDistance(t *testing.T) {
	sut := NewSimpleDistanceCalculator()
	result := sut.CalculateDistance(core.NewGene(1, "any1"), core.NewGene(2, "any2"))
	assert.NotNil(t, result)
	assert.IsType(t, float64(0), result)
}
