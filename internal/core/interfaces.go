package core

type (
	DistanceCalculator interface {
		CalculateDistance(from, to *Gene) float64
	}

	Selection interface {
		Select(*Population) *Chromosome
	}

	Crossover interface {
		Run() [2]*Chromosome
	}
)
