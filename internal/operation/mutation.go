package operation

import (
	"math/rand"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type Mutation struct{}

func NewMutation() *Mutation { return &Mutation{} }

func (m *Mutation) Mutate(chromosome *core.Chromosome, rate float64) {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)

	for i := range chromosome.Genes {
		if rnd.Float64() < rate {
			swapIndex := rnd.Intn(len(chromosome.Genes))
			chromosome.Genes[i], chromosome.Genes[swapIndex] = chromosome.Genes[swapIndex], chromosome.Genes[i]
		}
	}
}
