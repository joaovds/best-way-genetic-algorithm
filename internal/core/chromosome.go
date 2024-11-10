package core

type Chromosome struct {
	Genes         []Gene
	StartingPoint Gene
}

func NewChromosome(startingPoint Gene, genes []Gene) *Chromosome {
	return &Chromosome{
		Genes:         genes,
		StartingPoint: startingPoint,
	}
}
