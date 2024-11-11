package api

import "github.com/joaovds/best-way-genetic-algorithm/internal/core"

type (
	Location struct {
		Address    string `json:"address"`
		IsStarting bool   `json:"is_starting"`
	}

	LocationRequest []Location
)

func (l LocationRequest) ToCoreLocation() (startingPoint *core.Location, locations []*core.Location) {
	locations = make([]*core.Location, 0, len(l)-1)
	for i, location := range l {
		id := i + 1
		if location.IsStarting {
			startingPoint = core.NewLocation(id, location.Address)
		} else {
			locations = append(locations, core.NewLocation(id, location.Address))
		}
	}
	return
}
