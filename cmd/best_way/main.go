package main

import (
	"github.com/joaovds/best-way-genetic-algorithm/internal/algorithm"
	"github.com/joaovds/best-way-genetic-algorithm/internal/api"
)

func main() {
	locations := generateMockLocations()
	algorithmInstance := algorithm.NewAlgorithm(locations.ToCoreLocation())
	algorithmInstance.Run()
}

func generateMockLocations() api.LocationRequest {
	return api.LocationRequest{
		api.Location{Address: "any_address_1", IsStarting: false},
		api.Location{Address: "any_address_2", IsStarting: false},
		api.Location{Address: "any_address_3", IsStarting: true},
		api.Location{Address: "any_address_4", IsStarting: false},
	}
}
