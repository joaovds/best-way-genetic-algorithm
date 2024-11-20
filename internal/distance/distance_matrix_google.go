package distance

import (
	"context"
	"log"
	"sync"

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
		Origins:       locationsAddressStrings,
		Destinations:  locationsAddressStrings,
		DepartureTime: "now",
	}
	response, err := c.DistanceMatrix(ctx, req)
	if err != nil {
		log.Fatalf("error calling maps api: %s", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(response.Rows))
	for i, origin := range response.Rows {
		go func(i int) {
			defer wg.Done()
			for j, destination := range origin.Elements {
				originID := locationsAddressOrder[i]
				destinationID := locationsAddressOrder[j]
				if originID != destinationID {
					cache.CacheDistance(originID, destinationID, float64(destination.Distance.Meters), int(destination.DurationInTraffic.Seconds()))
				}
			}
		}(i)
	}
	wg.Wait()
}
