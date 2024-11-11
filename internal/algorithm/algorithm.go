package algorithm

import (
	"fmt"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
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
	fmt.Println(a.startingPoint)
	fmt.Println("Locales:")
	for _, localion := range a.locations {
		fmt.Println(localion)
	}
}
