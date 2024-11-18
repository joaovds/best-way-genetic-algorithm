package core

type (
	DistanceCalculator interface {
		CalculateDistance(from, to *Gene) float64
	}

	Selection interface {
		Select(*Population) *Chromosome
	}

	Crossover interface {
		Run(parent1, parent2 *Chromosome) [2]*Chromosome
	}

	Mutation interface {
		Mutate(chromosome *Chromosome, rate float64)
	}
)
