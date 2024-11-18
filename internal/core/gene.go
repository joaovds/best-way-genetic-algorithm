package core

type (
	Gene struct {
		Address  string
		Distance float64
		id       int
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

func (g *Gene) CalculateDistanceToDestination(destination *Gene, cache *Cache) float64 {
	if distance, exists := cache.GetFromCache(g.GetID(), destination.GetID()); exists {
		return distance
	}
	return 0
}
