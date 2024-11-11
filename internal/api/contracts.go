package api

import (
	"errors"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type (
	Location struct {
		Address    string `json:"address"`
		IsStarting bool   `json:"is_starting"`
	}

	LocationRequest []Location
)

func (l LocationRequest) ToCoreLocation() (startingPoint *core.Location, locations []*core.Location, err error) {
	if len(l) < 3 {
		return nil, nil, errors.New("invalid number of locations: must be at least 3")
	}

	startingPointFound := false
	locations = make([]*core.Location, 0, len(l)-1)
	for i, location := range l {
		id := i + 1
		if location.IsStarting && !startingPointFound {
			startingPoint = core.NewLocation(id, location.Address)
			startingPointFound = true
		} else {
			locations = append(locations, core.NewLocation(id, location.Address))
		}
	}

	if !startingPointFound {
		startingPoint = locations[0]
		locations = locations[1:]
	}

	return
}
