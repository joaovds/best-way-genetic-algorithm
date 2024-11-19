package core

type (
	Selection interface {
		Select(*Population) *Chromosome
	}

	Crossover interface {
		Run(parent1, parent2 *Chromosome) [2]*Chromosome
	}

	Mutation interface {
		Mutate(chromosome *Chromosome, rate float64)
	}

	DistanceCalculator interface {
		CalculateDistances(locations []*Location, cache *Cache)
	}
)
