package core

import (
	"math/rand/v2"
)

type Chromosome struct {
	StartingPoint *Gene
	Genes         []*Gene
	Fitness       float64
	TotalDistance float64
	SurvivalCount int
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

func (c *Chromosome) CalculateFitness(cache *Cache) float64 {
	distance := c.StartingPoint.CalculateDistanceToDestination(c.Genes[0], cache)
	c.StartingPoint.SetDistance(distance)
	fitness := distance

	for i, gene := range c.Genes {
		var distance float64
		if i == len(c.Genes)-1 {
			distance = gene.CalculateDistanceToDestination(c.StartingPoint, cache)
		} else {
			distance = gene.CalculateDistanceToDestination(c.Genes[i+1], cache)
		}
		gene.SetDistance(distance)
		fitness += distance
	}

	c.TotalDistance = fitness
	c.Fitness = 1 / fitness
	return c.Fitness
}
