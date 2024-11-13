package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulation_GetSize(t *testing.T) {
	population := NewPopulation(make([]*Chromosome, 3))
	assert.Equal(t, 3, population.GetSize())
}
