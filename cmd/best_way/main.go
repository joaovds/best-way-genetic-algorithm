package main

import (
	"github.com/joaovds/best-way-genetic-algorithm/internal/algorithm"
	"github.com/joaovds/best-way-genetic-algorithm/internal/dto"
)

func main() {
	locations := genMockLocations()

	algorithmGen := algorithm.NewAlgorithm(locations)

	algorithmGen.Run()
}

func genMockLocations() []*dto.Location {
	locations := []*dto.Location{
		dto.NewLocation(0.0, 1.0, false),
		dto.NewLocation(1.0, 3.0, false),
		dto.NewLocation(4.0, 4.0, true),
		dto.NewLocation(6.0, 1.0, false),
	}

	for i, location := range locations {
		location.ID = i
	}

	return locations
}
