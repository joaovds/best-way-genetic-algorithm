package core

type (
	Gene struct {
		Address  string
		Distance float64
		Duration int
		id       int
	}
)

func (g *Gene) GetID() int { return g.id }

func (g *Gene) SetDistance(distance float64) { g.Distance = distance }
func (g *Gene) SetDuration(duration int)     { g.Duration = duration }

func NewGene(id int, address string) *Gene {
	return &Gene{
		id:       id,
		Address:  address,
		Distance: 0.0,
		Duration: 0,
	}
}

func (g *Gene) CalculateDistanceToDestination(destination *Gene, cache *Cache) (distance float64, duration int) {
	if distance, duration, exists := cache.GetFromCache(g.GetID(), destination.GetID()); exists {
		return distance, duration
	}
	return 0, 0
}
