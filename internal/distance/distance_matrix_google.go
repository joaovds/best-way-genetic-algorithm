package distance

import (
	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type DistanceMatrixGoogle struct {
	API_KEY string
}

func NewDistanceMatrixGoogle(apiKey string) *DistanceMatrixGoogle {
	return &DistanceMatrixGoogle{
		API_KEY: apiKey,
	}
}

func (d *DistanceMatrixGoogle) CalculateDistances(locations []*core.Location, cache *core.Cache) {
	// var wg sync.WaitGroup
	// wg.Add(len(locations))
	// for i := range locations {
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		randSource := rand.NewSource(time.Now().UnixNano())
	// 		rnd := rand.New(randSource)
	//
	// 		for j := range locations {
	// 			if locations[i].ID != locations[j].ID {
	// 				cache.CacheDistance(locations[i].ID, locations[j].ID, rnd.Float64()*100)
	// 			}
	// 		}
	// 	}(i)
	// }
	// wg.Wait()
}
