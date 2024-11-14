package algorithm

import (
	"fmt"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/joaovds/best-way-genetic-algorithm/internal/distance"
)

const MAX_POPULATION_SIZE = 7000

type Algorithm struct {
	startingPoint  *core.Location
	locations      []*core.Location
	populationSize int
	chromosomeSize int
	maxGenerations int
}

func NewAlgorithm(startingPoint *core.Location, locations []*core.Location) *Algorithm {
	chromosomeSize := len(locations)
	populationSize := 1

	if chromosomeSize <= 5 {
		for i := range chromosomeSize {
			populationSize = populationSize * (i + 1)
		}
	} else {
		populationSize = chromosomeSize * 100
	}

	if populationSize > MAX_POPULATION_SIZE {
		populationSize = MAX_POPULATION_SIZE
	}

	return &Algorithm{
		startingPoint:  startingPoint,
		locations:      locations,
		populationSize: populationSize,
		chromosomeSize: chromosomeSize,
	}
}

func (a *Algorithm) Run() {
	distanceCalculator := distance.NewSimpleDistanceCalculator()
	population := core.GenerateInitialPopulation(a.populationSize, a.startingPoint, a.locations, core.GetCacheInstance)

	population.EvaluateFitness(distanceCalculator)

	fmt.Println("Location: ", a.startingPoint)
	fmt.Println("Locales:")
	for _, localion := range a.locations {
		fmt.Println(localion)
	}

	for _, c := range population.Chromosomes {
		fmt.Println("\n----- ... ----- \nStart:", c.StartingPoint)
		fmt.Println("\nGenes:")
		for _, gene := range c.Genes {
			fmt.Println(gene)
		}
	}
	fmt.Println("\nPopulation Size:", population.GetSize())
}
