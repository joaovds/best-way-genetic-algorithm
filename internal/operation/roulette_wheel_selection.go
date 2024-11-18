package operation

import (
	"math/rand"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type RouletteWheelSelection struct{}

func NewRouletteWheelSelection() *RouletteWheelSelection { return &RouletteWheelSelection{} }

func (r *RouletteWheelSelection) Select(population *core.Population) *core.Chromosome {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	nRandom := rnd.Float64() * population.TotalFitness
	sum := 0.0
	for _, chromosome := range population.Chromosomes {
		sum += chromosome.Fitness
		if sum >= nRandom {
			return chromosome
		}
	}
	return population.Chromosomes[len(population.Chromosomes)-1]
}
