package algorithm

type Chromosome struct {
	Genes   []*Gene
	Fitness int
}

func NewChromosome(genes []*Gene) *Chromosome {
	chromosome := new(Chromosome)
	chromosome.Genes = genes
	return chromosome
}
