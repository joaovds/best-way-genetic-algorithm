package core

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
