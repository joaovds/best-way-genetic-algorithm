package operation

import (
	"testing"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestRouletteWheelSelection(t *testing.T) {
	chromosome1 := core.NewChromosome(core.NewGene(1, "any"), []*core.Gene{core.NewGene(2, "any2")})
	chromosome2 := core.NewChromosome(core.NewGene(1, "any"), []*core.Gene{core.NewGene(3, "any3")})
	chromosome3 := core.NewChromosome(core.NewGene(1, "any"), []*core.Gene{core.NewGene(4, "any4")})

	chromosome1.Fitness = 10.0
	chromosome2.Fitness = 20.0
	chromosome3.Fitness = 70.0

	population := core.NewPopulation([]*core.Chromosome{chromosome1, chromosome2, chromosome3}, core.MockGetCacheInstanceFn, 4, 0.1)
	population.TotalFitness = chromosome1.Fitness + chromosome2.Fitness + chromosome3.Fitness

	selectionCounts := make(map[*core.Chromosome]int)
	rws := RouletteWheelSelection{}

	for i := 0; i < 10000; i++ {
		selected := rws.Select(population)
		selectionCounts[selected]++
	}

	assert.Greater(t, selectionCounts[chromosome3], selectionCounts[chromosome2])
	assert.Greater(t, selectionCounts[chromosome3], selectionCounts[chromosome1])
	assert.Greater(t, selectionCounts[chromosome2], selectionCounts[chromosome1])
}
