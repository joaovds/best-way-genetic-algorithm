package algorithm

import (
	"fmt"

	"github.com/joaovds/best-way-genetic-algorithm/internal/dto"
)

type (
	Algorithm struct {
		locations      []*dto.Location
		geneQuantity   int
		populationSize int
		maxGenerations int
	}
)

func NewAlgorithm(locations []*dto.Location) *Algorithm {
	return &Algorithm{
		locations:      locations,
		geneQuantity:   len(locations) - 1,
		populationSize: len(locations) * 10,
		maxGenerations: 300,
	}
}

func (a *Algorithm) Run() {
	population := a.InitPopulationWithLocations()

	fmt.Println(a.locations)
	println()
	for i, c := range population.Chromosomes {
		fmt.Println("Chromosome:", i, "=>", c.Fitness)
		fmt.Println("ID Starting:", c.StartingPoint.ID, "=> X:", c.StartingPoint.X, "Y:", c.StartingPoint.Y, "Next Point Distance:", c.StartingPoint.Distance)
		for _, g := range c.Genes {
			fmt.Println("ID:", g.ID, "=> X:", g.X, "Y:", g.Y, "Next Point Distance:", g.Distance)
		}
	}
}

func (a *Algorithm) InitPopulationWithLocations() *Population {
	chromosomes := make([]*Chromosome, a.populationSize)

	for index := range a.populationSize {
		genes := make([]*Gene, a.geneQuantity)
		var startingPointGene *Gene

		geneIndex := 0
		for _, location := range a.locations {
			if geneIndex == a.geneQuantity {
				startingPointGene = NewGene(location.ID, location.X, location.Y)
				continue
			}

			if location.StartingPoint {
				startingPointGene = NewGene(location.ID, location.X, location.Y)
			} else {
				genes[geneIndex] = NewGene(location.ID, location.X, location.Y)
				geneIndex++
			}
		}
		chromosomes[index] = NewChromosome(startingPointGene, genes)
		chromosomes[index].ShufflingGenes()
	}

	return NewPopulation(chromosomes)
}
