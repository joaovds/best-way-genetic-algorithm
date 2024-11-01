package algorithm

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
