package operation

import (
	"testing"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestMutation(t *testing.T) {
	chromosome := core.NewChromosome(core.NewGene(1, "any"), []*core.Gene{
		core.NewGene(2, "any2"),
		core.NewGene(3, "any3"),
		core.NewGene(4, "any4"),
		core.NewGene(5, "any5"),
	})

	initialGenes := append([]*core.Gene(nil), chromosome.Genes...)

	mutation := NewMutation()
	mutation.Mutate(chromosome, 0.999)
	assert.NotEqual(t, initialGenes, chromosome.Genes)
}
