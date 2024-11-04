package algorithm

import (
	"fmt"
	"math/rand"

	"github.com/joaovds/best-way-genetic-algorithm/internal/dto"
)

type (
	Algorithm struct {
		config              *Config
		locations           []*dto.Location
		geneQuantity        int
		numberOfChromosomes int
	}
)

func NewAlgorithm(config *Config, locations []*dto.Location) *Algorithm {
	return &Algorithm{
		config:              config,
		locations:           locations,
		geneQuantity:        len(locations),
		numberOfChromosomes: len(locations) * 10,
	}
}

func (a *Algorithm) Run() {
	fmt.Println(a.locations)
	println()
	for i, c := range a.InitPopulationWithLocations().Chromosomes {
		fmt.Println("Chromosome:", i, "=>", c.Fitness)
		for _, g := range c.Genes {
			fmt.Println("ID:", g.ID, "=> X:", g.X, "Y:", g.Y)
		}
	}
}

func (a *Algorithm) InitPopulationWithLocations() *Population {
	chromosomes := make([]*Chromosome, a.numberOfChromosomes)

	for index := range a.numberOfChromosomes {
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
