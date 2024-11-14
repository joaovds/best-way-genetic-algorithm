package core

import "sync"

type Population struct {
	Chromosomes []*Chromosome
}

func (p *Population) GetSize() int { return len(p.Chromosomes) }

func NewPopulation(chromosomes []*Chromosome) *Population {
	return &Population{chromosomes}
}

func GenerateInitialPopulation(size int, startingPoint *Location, locations []*Location) *Population {
	chromosomes := make([]*Chromosome, size)
	for i := range size {
		chromosomes[i] = NewChromosome(startingPoint.ToNewGene(), LocationsToGenes(locations))
		chromosomes[i].ShufflingGenes()
	}
	return NewPopulation(chromosomes)
}

func (p *Population) EvaluateFitness(dc DistanceCalculator) {
	var wg sync.WaitGroup

	for _, chromosome := range p.Chromosomes {
		wg.Add(1)
		go func(c *Chromosome) {
			defer wg.Done()
			c.CalculateFitness(dc)
		}(chromosome)
	}

	wg.Wait()
}
