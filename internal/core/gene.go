package core

import (
	"fmt"
	"sync"
)

var (
	distancesCache map[string]float64 = make(map[string]float64)
	cacheMutex     sync.Mutex
	cleanCache     = func() {
		cacheMutex.Lock()
		defer cacheMutex.Unlock()
		distancesCache = make(map[string]float64)
	}
)

type (
	Gene struct {
		Address  string
		Distance float64
		id       int
	}

	DistanceCalculator interface {
		CalculateDistance(from, to *Gene) float64
	}
)

func (g *Gene) GetID() int { return g.id }

func (g *Gene) SetDistance(distance float64) { g.Distance = distance }

func NewGene(id int, address string) *Gene {
	return &Gene{
		id:       id,
		Address:  address,
		Distance: 0.0,
	}
}

func (g *Gene) CalculateDistanceToDestination(destination *Gene, calculator DistanceCalculator) float64 {
	if g.id == destination.id {
		return 0
	}

	cacheKey := generateCacheKey(g.id, destination.id)
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	if distance, exists := distancesCache[cacheKey]; exists {
		return distance
	}

	distance := calculator.CalculateDistance(g, destination)
	distancesCache[cacheKey] = distance
	return distance
}

func generateCacheKey(fromID, destinationID int) string {
	return fmt.Sprintf("%d-%d", fromID, destinationID)
}
