package core

type Gene struct {
	Address       string
	Distance      float64
	id            int
	startingPoint bool
}

func (g *Gene) GetID() int { return g.id }

func (g *Gene) IsStartingPoint() bool { return g.startingPoint }

func (g *Gene) SetStartingPoint() { g.startingPoint = true }

func NewGene(id int, address string) *Gene {
	return &Gene{
		id:            id,
		Address:       address,
		Distance:      0.0,
		startingPoint: false,
	}
}
