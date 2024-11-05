package algorithm

import "math/rand/v2"

type Chromosome struct {
	StartingPoint *Gene
	Genes         []*Gene
	Fitness       int
}

func NewChromosome(startingPoint *Gene, genes []*Gene) *Chromosome {
	chromosome := new(Chromosome)
	chromosome.StartingPoint = startingPoint
	chromosome.Genes = genes
	return chromosome
}

func (c *Chromosome) ShufflingGenes() {
	shuffledGenesOrder := rand.Perm(len(c.Genes))
	shuffledGenes := make([]*Gene, len(c.Genes))
	for i, newIndex := range shuffledGenesOrder {
		shuffledGenes[i] = c.Genes[newIndex]
	}
	c.Genes = shuffledGenes
}
