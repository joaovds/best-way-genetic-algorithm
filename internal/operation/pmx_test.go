package operation

import (
	"testing"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestPMX(t *testing.T) {
	t.Run("should return different chromosome size error", func(t *testing.T) {
		parent1Genes := []*core.Gene{
			core.NewGene(1, "any"),
			core.NewGene(2, "any"),
			core.NewGene(3, "any"),
		}
		parent2Genes := []*core.Gene{
			core.NewGene(2, "any"),
			core.NewGene(3, "any"),
		}
		parent1 := core.NewChromosome(nil, parent1Genes)
		parent2 := core.NewChromosome(nil, parent2Genes)

		assert.PanicsWithError(t, "The chromosomes must contain the same number of genes", func() {
			NewPMX().Run(parent1, parent2)
		})
	})

	t.Run("should return the crossed children with the correct values", func(t *testing.T) {
		t.Run("1", func(t *testing.T) {
			parent1Genes := []*core.Gene{
				core.NewGene(1, "any"),
				core.NewGene(2, "any"),
				core.NewGene(3, "any"),
				core.NewGene(4, "any"),
				core.NewGene(5, "any"),
				core.NewGene(6, "any"),
				core.NewGene(7, "any"),
				core.NewGene(8, "any"),
				core.NewGene(9, "any"),
			}
			parent2Genes := []*core.Gene{
				core.NewGene(5, "any"),
				core.NewGene(4, "any"),
				core.NewGene(6, "any"),
				core.NewGene(9, "any"),
				core.NewGene(2, "any"),
				core.NewGene(1, "any"),
				core.NewGene(7, "any"),
				core.NewGene(8, "any"),
				core.NewGene(3, "any"),
			}
			parent1 := core.NewChromosome(core.NewGene(100, "starting"), parent1Genes)
			parent2 := core.NewChromosome(core.NewGene(100, "starting"), parent2Genes)
			pmx := NewPMX()
			// mock
			pmx.StartPoint, pmx.EndPoint = 2, 5

			children := pmx.Run(parent1, parent2)
			child1 := children[0]
			child2 := children[1]

			t.Run("child1 values", func(t *testing.T) {
				assert.Equal(t, 3, child1.Genes[0].GetID())
				assert.Equal(t, 5, child1.Genes[1].GetID())
				assert.Equal(t, 6, child1.Genes[2].GetID())
				assert.Equal(t, 9, child1.Genes[3].GetID())
				assert.Equal(t, 2, child1.Genes[4].GetID())
				assert.Equal(t, 1, child1.Genes[5].GetID())
				assert.Equal(t, 7, child1.Genes[6].GetID())
				assert.Equal(t, 8, child1.Genes[7].GetID())
				assert.Equal(t, 4, child1.Genes[8].GetID())
			})

			t.Run("child2 values", func(t *testing.T) {
				assert.Equal(t, 2, child2.Genes[0].GetID())
				assert.Equal(t, 9, child2.Genes[1].GetID())
				assert.Equal(t, 3, child2.Genes[2].GetID())
				assert.Equal(t, 4, child2.Genes[3].GetID())
				assert.Equal(t, 5, child2.Genes[4].GetID())
				assert.Equal(t, 6, child2.Genes[5].GetID())
				assert.Equal(t, 7, child2.Genes[6].GetID())
				assert.Equal(t, 8, child2.Genes[7].GetID())
				assert.Equal(t, 1, child2.Genes[8].GetID())
			})
		})

		t.Run("2", func(t *testing.T) {
			parent1Genes := []*core.Gene{
				core.NewGene(9, "any"),
				core.NewGene(2, "any"),
				core.NewGene(7, "any"),
				core.NewGene(5, "any"),
				core.NewGene(4, "any"),
				core.NewGene(3, "any"),
				core.NewGene(6, "any"),
				core.NewGene(1, "any"),
				core.NewGene(8, "any"),
			}
			parent2Genes := []*core.Gene{
				core.NewGene(2, "any"),
				core.NewGene(8, "any"),
				core.NewGene(3, "any"),
				core.NewGene(6, "any"),
				core.NewGene(9, "any"),
				core.NewGene(5, "any"),
				core.NewGene(7, "any"),
				core.NewGene(4, "any"),
				core.NewGene(1, "any"),
			}
			parent1 := core.NewChromosome(core.NewGene(100, "starting"), parent1Genes)
			parent2 := core.NewChromosome(core.NewGene(100, "starting"), parent2Genes)
			pmx := NewPMX()
			// mock
			pmx.StartPoint, pmx.EndPoint = 3, 5

			children := pmx.Run(parent1, parent2)
			child1 := children[0]
			child2 := children[1]

			t.Run("child1 values", func(t *testing.T) {
				assert.Equal(t, 4, child1.Genes[0].GetID())
				assert.Equal(t, 2, child1.Genes[1].GetID())
				assert.Equal(t, 7, child1.Genes[2].GetID())
				assert.Equal(t, 6, child1.Genes[3].GetID())
				assert.Equal(t, 9, child1.Genes[4].GetID())
				assert.Equal(t, 5, child1.Genes[5].GetID())
				assert.Equal(t, 3, child1.Genes[6].GetID())
				assert.Equal(t, 1, child1.Genes[7].GetID())
				assert.Equal(t, 8, child1.Genes[8].GetID())
			})

			t.Run("child2 values", func(t *testing.T) {
				assert.Equal(t, 2, child2.Genes[0].GetID())
				assert.Equal(t, 8, child2.Genes[1].GetID())
				assert.Equal(t, 6, child2.Genes[2].GetID())
				assert.Equal(t, 5, child2.Genes[3].GetID())
				assert.Equal(t, 4, child2.Genes[4].GetID())
				assert.Equal(t, 3, child2.Genes[5].GetID())
				assert.Equal(t, 7, child2.Genes[6].GetID())
				assert.Equal(t, 9, child2.Genes[7].GetID())
				assert.Equal(t, 1, child2.Genes[8].GetID())
			})
		})

		t.Run("3", func(t *testing.T) {
			parent1Genes := []*core.Gene{
				core.NewGene(1, "any"),
				core.NewGene(2, "any"),
				core.NewGene(3, "any"),
				core.NewGene(4, "any"),
			}
			parent2Genes := []*core.Gene{
				core.NewGene(2, "any"),
				core.NewGene(3, "any"),
				core.NewGene(1, "any"),
				core.NewGene(4, "any"),
			}
			parent1 := core.NewChromosome(core.NewGene(100, "starting"), parent1Genes)
			parent2 := core.NewChromosome(core.NewGene(100, "starting"), parent2Genes)
			pmx := NewPMX()
			// mock
			pmx.StartPoint, pmx.EndPoint = 0, 1

			children := pmx.Run(parent1, parent2)
			child1 := children[0]
			child2 := children[1]

			t.Run("child1 values", func(t *testing.T) {
				assert.Equal(t, 2, child1.Genes[0].GetID())
				assert.Equal(t, 3, child1.Genes[1].GetID())
				assert.Equal(t, 1, child1.Genes[2].GetID())
				assert.Equal(t, 4, child1.Genes[3].GetID())
			})

			t.Run("child2 values", func(t *testing.T) {
				assert.Equal(t, 1, child2.Genes[0].GetID())
				assert.Equal(t, 2, child2.Genes[1].GetID())
				assert.Equal(t, 3, child2.Genes[2].GetID())
				assert.Equal(t, 4, child2.Genes[3].GetID())
			})
		})
	})
}

func TestContains(t *testing.T) {
	genes := []*core.Gene{
		core.NewGene(1, "any"),
		core.NewGene(4, "any"),
		core.NewGene(16, "any"),
		core.NewGene(28, "any"),
		core.NewGene(10, "any"),
	}
	existsGene := core.NewGene(28, "any")
	noExistsGene := core.NewGene(282, "any")

	assert.True(t, contains(genes, existsGene))
	assert.False(t, contains(genes, noExistsGene))
}

func TestIndexOfGene(t *testing.T) {
	genes := []*core.Gene{
		core.NewGene(1, "any"),
		core.NewGene(4, "any"),
		core.NewGene(16, "any"),
		core.NewGene(28, "any"),
		core.NewGene(10, "any"),
	}

	assert.Equal(t, 3, indexOfGene(genes, core.NewGene(28, "any")))
	assert.Equal(t, 1, indexOfGene(genes, core.NewGene(4, "any")))
}

func TestGetCrossoverPoints(t *testing.T) {
	t.Run("ValidCrossoverPoints", func(t *testing.T) {
		size := 10
		for i := 0; i < 1000; i++ {
			startPoint, endPoint := getCrossoverPoints(size)

			assert.GreaterOrEqual(t, startPoint, 0)
			assert.Less(t, startPoint, size)

			assert.GreaterOrEqual(t, endPoint, 0)
			assert.Less(t, endPoint, size)
			assert.Less(t, startPoint, endPoint)
		}
	})

	t.Run("EdgeCases", func(t *testing.T) {
		t.Run("Size2", func(t *testing.T) {
			startPoint, endPoint := getCrossoverPoints(2)
			assert.Equal(t, startPoint, 0)
			assert.Equal(t, endPoint, 1)
		})

		t.Run("Size1", func(t *testing.T) {
			assert.PanicsWithError(t, "The number must be greater than 1", func() {
				getCrossoverPoints(1)
			})
		})

		t.Run("Size0", func(t *testing.T) {
			assert.PanicsWithError(t, "The number must be greater than 1", func() {
				getCrossoverPoints(0)
			})
		})
	})
}
