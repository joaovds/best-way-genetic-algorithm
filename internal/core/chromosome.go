package core

import (
	"math/rand/v2"
)

type Chromosome struct {
	Genes         []Gene
	StartingPoint Gene
	Fitness       float64
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

func (c *Chromosome) CalculateFitness() float64 {
	fitness := c.StartingPoint.Distance
	for _, gene := range c.Genes {
		fitness += gene.Distance
	}
	c.Fitness = fitness
	return fitness
}
