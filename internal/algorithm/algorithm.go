package algorithm

import (
	"fmt"

	"github.com/joaovds/best-way-genetic-algorithm/internal/dto"
)

type (
	Algorithm struct {
		config    *Config
		locations []*dto.Location
	}
)

func NewAlgorithm(config *Config, locations []*dto.Location) *Algorithm {
	return &Algorithm{config: config, locations: locations}
}

func (a *Algorithm) Run() {
	fmt.Println(a.locations)
}
