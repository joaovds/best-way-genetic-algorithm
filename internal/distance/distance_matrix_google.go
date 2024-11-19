package distance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"googlemaps.github.io/maps"
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
	c, err := maps.NewClient(maps.WithAPIKey(d.API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	locationsAddressStrings := make([]string, 0, len(locations))
	locationsAddressOrder := make(map[int]int)
	for i, loc := range locations {
		locationsAddressStrings = append(locationsAddressStrings, loc.Address)
		locationsAddressOrder[i] = loc.ID
	}

	ctx := context.Background()
	req := &maps.DistanceMatrixRequest{
		Origins:      locationsAddressStrings,
		Destinations: locationsAddressStrings,
	}
	response, err := c.DistanceMatrix(ctx, req)
	if err != nil {
		log.Fatalf("error calling maps api: %s", err)
	}

	responseJson, _ := json.Marshal(response)
	fmt.Println(string(responseJson))
	fmt.Println()
	fmt.Println(locationsAddressStrings)
	fmt.Println(locationsAddressOrder)

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
