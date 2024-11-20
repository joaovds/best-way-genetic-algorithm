package core

import (
	"math/rand/v2"
)

type Chromosome struct {
	StartingPoint *Gene
	Genes         []*Gene
	Fitness       float64
	TotalDistance float64
	TotalDuration int
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
	distance, duration := c.StartingPoint.CalculateDistanceToDestination(c.Genes[0], cache)
	c.StartingPoint.SetDistance(distance)
	c.StartingPoint.SetDuration(duration)
	fitness := distance
	totalDuration := duration

	for i, gene := range c.Genes {
		var distance float64
		var duration int
		if i == len(c.Genes)-1 {
			distance, duration = gene.CalculateDistanceToDestination(c.StartingPoint, cache)
		} else {
			distance, duration = gene.CalculateDistanceToDestination(c.Genes[i+1], cache)
		}
		gene.SetDistance(distance)
		gene.SetDuration(duration)
		fitness += distance
		totalDuration += duration
	}

	c.TotalDistance = fitness
	c.TotalDuration = totalDuration

	c.Fitness = 1 / fitness
	return c.Fitness
}
