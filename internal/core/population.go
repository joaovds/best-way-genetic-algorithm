package core

import "sync"

type Population struct {
	Chromosomes  []*Chromosome
	TotalFitness float64
}

func (p *Population) GetSize() int { return len(p.Chromosomes) }

func NewPopulation(chromosomes []*Chromosome) *Population {
	return &Population{Chromosomes: chromosomes}
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
	totalFitnessCh := make(chan float64)

	for _, chromosome := range p.Chromosomes {
		wg.Add(1)
		go func(c *Chromosome) {
			defer wg.Done()
			totalFitnessCh <- c.CalculateFitness(dc)
		}(chromosome)
	}

	go func() {
		wg.Wait()
		close(totalFitnessCh)
	}()

	for fitness := range totalFitnessCh {
		p.TotalFitness += fitness
	}
}
