package core

import "math/rand/v2"

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

func (c *Chromosome) ShufflingGenes() {
	shuffledGenesOrder := rand.Perm(len(c.Genes))
	shuffledGenes := make([]Gene, len(c.Genes))
	for i, newIndex := range shuffledGenesOrder {
		shuffledGenes[i] = c.Genes[newIndex]
	}
	c.Genes = shuffledGenes
}
