package algorithm

import (
	"fmt"
	"math/rand"

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
		geneQuantity:   len(locations),
		populationSize: len(locations) * 10,
		maxGenerations: 300,
	}
}

func (a *Algorithm) Run() {
	fmt.Println(a.locations)
	println()
	for i, c := range a.InitPopulationWithLocations().Chromosomes {
		fmt.Println("Chromosome:", i, "=>", c.Fitness)
		for _, g := range c.Genes {
			fmt.Println("ID:", g.ID, "=> X:", g.X, "Y:", g.Y, "Next Point Distance:", g.Distance)
		}
	}
}

func (a *Algorithm) InitPopulationWithLocations() *Population {
	chromosomes := make([]*Chromosome, a.populationSize)

	for index := range a.populationSize {
		shuffledLocations := rand.Perm(len(a.locations))
		genes := make([]*Gene, a.geneQuantity)

		for index, locationIndex := range shuffledLocations {
			location := a.locations[locationIndex]
			genes[index] = NewGene(location.ID, location.X, location.Y)
		}
		chromosomes[index] = NewChromosome(genes)
	}

	return NewPopulation(chromosomes)
}
