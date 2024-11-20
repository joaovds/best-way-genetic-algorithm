package distance

import (
	"math/rand"
	"sync"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type InBatchCalculator struct{}

func NewInBatchCalculator() *InBatchCalculator { return new(InBatchCalculator) }

func (i *InBatchCalculator) CalculateDistances(locations []*core.Location, cache *core.Cache) {
	var wg sync.WaitGroup
	wg.Add(len(locations))
	for i := range locations {
		go func(i int) {
			defer wg.Done()
			randSource := rand.NewSource(time.Now().UnixNano())
			rnd := rand.New(randSource)

			for j := range locations {
				if locations[i].ID != locations[j].ID {
					cache.CacheDistance(locations[i].ID, locations[j].ID, rnd.Float64()*100, rnd.Intn(10000))
				}
			}
		}(i)
	}
	wg.Wait()
}
