package algorithm

type Population struct {
	Chromosomes []*Chromosome
	Size        int
}

func NewPopulation(chromosome []*Chromosome) *Population {
	return &Population{
		Chromosomes: chromosome,
		Size:        len(chromosome),
	}
}
