package main

import (
	"github.com/joaovds/best-way-genetic-algorithm/internal/algorithm"
	"github.com/joaovds/best-way-genetic-algorithm/internal/api"
)

func main() {
	locations := generateMockLocations()
	startingPoint, coreLocations, err := locations.ToCoreLocation()
	if err != nil {
		panic(err)
	}
	config := algorithm.NewConfig(7000, 300, 4, 0.4)
	algorithmInstance := algorithm.NewAlgorithm(config, startingPoint, coreLocations)
	algorithmInstance.Run()
	algorithmInstance.RenderChart()
}

func generateMockLocations() api.LocationRequest {
	return api.LocationRequest{
		api.Location{Address: "any_address_1", IsStarting: false},
		api.Location{Address: "any_address_2", IsStarting: false},
		api.Location{Address: "any_address_3", IsStarting: true},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
		api.Location{Address: "any_address_4", IsStarting: false},
	}
}
