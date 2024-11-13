package core

type Population struct {
	Chromosomes []*Chromosome
}

func (p *Population) GetSize() int { return len(p.Chromosomes) }

func NewPopulation(chromosomes []*Chromosome) *Population {
	return &Population{chromosomes}
}
