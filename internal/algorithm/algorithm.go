package algorithm

import (
	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/joaovds/best-way-genetic-algorithm/internal/operation"
)

type (
	Algorithm struct {
		distanceCalculator core.DistanceCalculator
		config             *Config
		startingPoint      *core.Location
		locations          []*core.Location
		stats              []generationStats
		populationSize     int
		chromosomeSize     int
	}

	AlgorithmResponse struct {
		BestWay        *core.Chromosome
		PopulationSize int
		MaxGenerations int
		ElitismNumber  int
		MutationRate   float64
	}

	generationStats struct {
		better stats
		middle stats
		worse  stats
	}

	stats struct {
		distance, fitness float64
	}
)

func NewAlgorithm(config *Config, startingPoint *core.Location, locations []*core.Location, distanceCalculator core.DistanceCalculator) *Algorithm {
	chromosomeSize := len(locations)
	populationSize := 1

	if chromosomeSize <= 5 {
		for i := range chromosomeSize {
			populationSize = populationSize * (i + 1)
		}
	} else if chromosomeSize <= 12 {
		populationSize = chromosomeSize * 10
	} else if chromosomeSize <= 32 {
		populationSize = chromosomeSize * 40
	} else {
		populationSize = chromosomeSize * 60
	}

	if populationSize > config.MaxPopulationSize {
		populationSize = config.MaxPopulationSize
	}

	return &Algorithm{
		config:             config,
		startingPoint:      startingPoint,
		locations:          locations,
		populationSize:     populationSize,
		chromosomeSize:     chromosomeSize,
		distanceCalculator: distanceCalculator,
	}
}

func (a *Algorithm) Run() *AlgorithmResponse {
	core.GetCacheInstance().Clean()
	a.distanceCalculator.CalculateDistances(append(a.locations, a.startingPoint), core.GetCacheInstance())
	selection := operation.NewRouletteWheelSelection()
	crossover := operation.NewPMX()
	mutation := operation.NewMutation()

	population := core.GenerateInitialPopulation(
		a.populationSize,
		a.startingPoint,
		a.locations,
		core.GetCacheInstance,
		a.config.ElitismNumber,
		a.config.MutationRate,
	)

	for range a.config.MaxGenerations {
		population.EvaluateFitness()
		population.SortByFitness()

		a.stats = append(a.stats, generationStats{
			better: stats{
				distance: population.Chromosomes[0].TotalDistance,
				fitness:  population.Chromosomes[0].Fitness,
			},
			middle: stats{
				distance: population.Chromosomes[population.GetSize()/2].TotalDistance,
				fitness:  population.Chromosomes[population.GetSize()/2].Fitness,
			},
			worse: stats{
				distance: population.Chromosomes[population.GetSize()-1].TotalDistance,
				fitness:  population.Chromosomes[population.GetSize()-1].Fitness,
			},
		})

		population = population.GenerateNextGeration(selection, crossover, mutation)
	}

	return &AlgorithmResponse{
		BestWay:        population.Chromosomes[0],
		PopulationSize: population.GetSize(),
		MaxGenerations: a.config.MaxGenerations,
		ElitismNumber:  a.config.ElitismNumber,
		MutationRate:   a.config.MutationRate,
	}
}
