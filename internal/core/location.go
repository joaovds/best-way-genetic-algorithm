package core

type Location struct {
	Address string
	ID      int
}

func NewLocation(id int, address string) *Location {
	return &Location{address, id}
}

func (l *Location) ToNewGene() Gene {
	return NewGene(l.ID, l.Address)
}
