package algorithm

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

var (
	distancesCache map[string]float64 = make(map[string]float64)
	cacheMutex     sync.Mutex
)

type Gene struct {
	ID            int
	X             float64
	Y             float64
	Distance      float64 // distance to the next point
	StartingPoint bool
}

func NewGene(id int, x, y float64) *Gene {
	return &Gene{
		ID:            id,
		X:             x,
		Y:             y,
		Distance:      0.0,
		StartingPoint: false,
	}
}

func (g *Gene) CalculateDistanceToDestination(destination *Gene) float64 {
	if g.ID == destination.ID {
		return 0
	}

	cacheKey := g.generateCacheKey(destination.ID)
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	if distance, exists := distancesCache[cacheKey]; exists {
		return distance
	}
	distance := rand.Float64() * 10
	distancesCache[cacheKey] = distance
	return distance
}

func (g *Gene) generateCacheKey(destinationID int) string {
	return fmt.Sprintf("%d-%d", g.ID, destinationID)
}
