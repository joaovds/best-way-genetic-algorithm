package algorithm

import "math/rand/v2"

type Gene struct {
	ID       int
	X        float64
	Y        float64
	Distance float64 // distance to the next point
}

func NewGene(id int, x, y float64) *Gene {
	return &Gene{
		ID:       id,
		X:        x,
		Y:        y,
		Distance: 0.0,
	}
}

func (g *Gene) CalculateDistanceToDestination(destination *Gene) float64 {
	if g.ID == destination.ID {
		return 0
	}

	return rand.Float64() * 10
}
