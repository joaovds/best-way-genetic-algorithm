package algorithm

type Population struct {
	Chromosomes []*Chromosome
	Size        int
}

func NewPopulation(chromosomes []*Chromosome) *Population {
	for _, chromosome := range chromosomes {
		chromosome.StartingPoint.Distance = chromosome.StartingPoint.CalculateDistanceToDestination(chromosome.Genes[0])

		for i, gene := range chromosome.Genes {
			if i == len(chromosome.Genes)-1 {
				gene.Distance = gene.CalculateDistanceToDestination(chromosome.StartingPoint)
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
