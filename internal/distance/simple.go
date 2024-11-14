package distance

import (
	"math/rand"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type SimpleDistanceCalculator struct{}

func NewSimpleDistanceCalculator() *SimpleDistanceCalculator { return new(SimpleDistanceCalculator) }

func (s *SimpleDistanceCalculator) CalculateDistance(from, to *core.Gene) float64 {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	return rnd.Float64() * 100
}
