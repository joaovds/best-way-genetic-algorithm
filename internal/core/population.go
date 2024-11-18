package core

import (
	"sort"
	"sync"
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
		return p.Chromosomes[i].Fitness > p.Chromosomes[j].Fitness
	})
}
