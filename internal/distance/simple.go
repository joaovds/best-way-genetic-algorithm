package distance

import (
	"math/rand"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

var Counter int

type SimpleDistanceCalculator struct{}

func NewSimpleDistanceCalculator() *SimpleDistanceCalculator { return new(SimpleDistanceCalculator) }

func (s *SimpleDistanceCalculator) CalculateDistance(from, to *core.Gene) float64 {
	Counter += 1
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	return rnd.Float64() * 100
}
