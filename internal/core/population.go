package core

import (
	"sort"
	"sync"
)

const (
	ELITISM_NUMBER = 4
	MUTATION_RATE  = 0.1
)

type Population struct {
	cache        *Cache
	Chromosomes  []*Chromosome
	TotalFitness float64
}

func (p *Population) GetSize() int { return len(p.Chromosomes) }

func NewPopulation(chromosomes []*Chromosome, getCacheInstanceFn GetCacheInstanceFn) *Population {
	return &Population{Chromosomes: chromosomes, cache: getCacheInstanceFn()}
}

func GenerateInitialPopulation(size int, startingPoint *Location, locations []*Location, getCacheInstanceFn GetCacheInstanceFn) *Population {
	chromosomes := make([]*Chromosome, size)
	for i := range size {
		chromosomes[i] = NewChromosome(startingPoint.ToNewGene(), LocationsToGenes(locations))
		chromosomes[i].ShufflingGenes()
	}
	return NewPopulation(chromosomes, getCacheInstanceFn)
}

func (p *Population) EvaluateFitness(dc DistanceCalculator) {
	var wg sync.WaitGroup
	totalFitnessCh := make(chan float64, len(p.Chromosomes))

	for _, chromosome := range p.Chromosomes {
		wg.Add(1)
		go func(c *Chromosome) {
			defer wg.Done()
			totalFitnessCh <- c.CalculateFitness(dc, p.cache)
		}(chromosome)
	}

	go func() {
		wg.Wait()
		close(totalFitnessCh)
	}()

	p.TotalFitness = 0
	for fitness := range totalFitnessCh {
		p.TotalFitness += fitness
	}
}

func (p *Population) SortByFitness() {
	sort.Slice(p.Chromosomes, func(i, j int) bool {
		if p.Chromosomes[i].Fitness == p.Chromosomes[j].Fitness {
			return p.Chromosomes[i].SurvivalCount > p.Chromosomes[j].SurvivalCount
		}

		return p.Chromosomes[i].Fitness > p.Chromosomes[j].Fitness
	})
}

func (p *Population) GenerateNextGeration(selection Selection, crossover Crossover, mutation Mutation) *Population {
	nextGenerationChromosomes := make([]*Chromosome, p.GetSize())
	p.Elitism(nextGenerationChromosomes)
	numberOfNewChromosomes := ELITISM_NUMBER

	for {
		parent1 := selection.Select(p)
		parent2 := selection.Select(p)
		children := crossover.Run(parent1, parent2)

		for i := range children {
			mutation.Mutate(children[i], MUTATION_RATE)
		}

		for i := range children {
			if numberOfNewChromosomes < len(nextGenerationChromosomes) {
				nextGenerationChromosomes[numberOfNewChromosomes] = children[i]
				numberOfNewChromosomes++
			}
		}

		if numberOfNewChromosomes >= len(nextGenerationChromosomes) {
			break
		}
	}

	return NewPopulation(nextGenerationChromosomes, GetCacheInstance)
}

func (p *Population) Elitism(nextGenerationChromosomes []*Chromosome) {
	for i := range ELITISM_NUMBER {
		if i == p.GetSize() {
			break
		}
		nextGenerationChromosomes[i] = p.Chromosomes[i]
		nextGenerationChromosomes[i].SurvivalCount++
	}
}
