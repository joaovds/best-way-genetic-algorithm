package core

import (
	"math/rand/v2"
)

type Chromosome struct {
	StartingPoint *Gene
	Genes         []*Gene
	Fitness       float64
}

func NewChromosome(startingPoint *Gene, genes []*Gene) *Chromosome {
	return &Chromosome{
		Genes:         genes,
		StartingPoint: startingPoint,
	}
}

func (c *Chromosome) ShufflingGenes() {
	shuffledGenesOrder := rand.Perm(len(c.Genes))
	shuffledGenes := make([]*Gene, len(c.Genes))
	for i, newIndex := range shuffledGenesOrder {
		shuffledGenes[i] = c.Genes[newIndex]
	}
	c.Genes = shuffledGenes
}

func (c *Chromosome) CalculateFitness(dc DistanceCalculator) float64 {
	distance := c.StartingPoint.CalculateDistanceToDestination(c.Genes[0], dc)
	c.StartingPoint.SetDistance(distance)
	fitness := distance

	for i, gene := range c.Genes {
		var distance float64
		if i == len(c.Genes)-1 {
			distance = gene.CalculateDistanceToDestination(c.StartingPoint, dc)
		} else {
			distance = gene.CalculateDistanceToDestination(c.Genes[i+1], dc)
		}
		gene.SetDistance(distance)
		fitness += distance
	}

	c.Fitness = fitness
	return fitness
}
