package core

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

func (g *Gene) CalculateDistanceToDestination(destination *Gene, calculator DistanceCalculator, cache *Cache) float64 {
	if g.id == destination.id {
		return 0
	}

	lock := getCacheLock(g.GetID(), destination.GetID())
	lock.Lock()
	defer lock.Unlock()

	if distance, exists := cache.GetFromCache(g.GetID(), destination.GetID()); exists {
		return distance
	}

	distance := calculator.CalculateDistance(g, destination)
	cache.CacheDistance(g.GetID(), destination.GetID(), distance)
	return distance
}
