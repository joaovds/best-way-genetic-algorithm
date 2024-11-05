package algorithm

type Population struct {
	Chromosomes []*Chromosome
	Size        int
}

func NewPopulation(chromosomes []*Chromosome) *Population {
	for _, chromosome := range chromosomes {
		for i, gene := range chromosome.Genes {
			if i == len(chromosome.Genes)-1 {
				gene.Distance = gene.CalculateDistanceToDestination(chromosome.Genes[0])
			} else {
				gene.Distance = gene.CalculateDistanceToDestination(chromosome.Genes[i+1])
			}
		}
	}

	return &Population{
		Chromosomes: chromosomes,
		Size:        len(chromosomes),
	}
}
