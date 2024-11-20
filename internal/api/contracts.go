package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/algorithm"
	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type (
	Location struct {
		Address    string `json:"address"`
		IsStarting bool   `json:"is_starting"`
	}

	LocationRequest struct {
		Locations []Location `json:"locations"`
	}

	LocationRes struct {
		Address               string  `json:"address"`
		DistanceHumanReadable string  `json:"distance_human_readable"`
		TimeHumanReadable     string  `json:"time_human_readable"`
		IsStarting            bool    `json:"is_starting"`
		DistanceToNextPoint   float64 `json:"distance_to_next_point"`
		TimeSeconds           int     `json:"time_in_seconds"`
	}

	Response struct {
		TotalDistanceHumanReadable string        `json:"total_distance_human_readable"`
		TotalTimeHumanReadable     string        `json:"total_time_human_readable"`
		Route                      []LocationRes `json:"route"`
		TotalTime                  int           `json:"total_time"`
		PopulationSize             int           `json:"population_size"`
		MaxGenerations             int           `json:"max_generations"`
		ElitismNumber              int           `json:"elitism_number"`
		TotalDistance              float64       `json:"total_distance"`
		MutationRate               float64       `json:"mutation_rate"`
	}
)

func (l *LocationRequest) Validate() error {
	if len(l.Locations) <= 2 {
		return errors.New("The locations must be more than 2")
	}
	return nil
}

func (l LocationRequest) ToCoreLocation() (startingPoint *core.Location, locations []*core.Location, err error) {
	if len(l.Locations) < 3 {
		return nil, nil, errors.New("invalid number of locations: must be at least 3")
	}

	startingPointFound := false
	locations = make([]*core.Location, 0, len(l.Locations)-1)
	for i, location := range l.Locations {
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

func AlgorithmResponseToApiResponse(algorithmRes *algorithm.AlgorithmResponse) Response {
	route := make([]LocationRes, 0, len(algorithmRes.BestWay.Genes)+1)
	route = append(route, LocationRes{
		Address:               algorithmRes.BestWay.StartingPoint.Address,
		IsStarting:            true,
		DistanceToNextPoint:   algorithmRes.BestWay.StartingPoint.Distance,
		DistanceHumanReadable: formatDistance(int(algorithmRes.BestWay.StartingPoint.Distance)),
		TimeSeconds:           algorithmRes.BestWay.StartingPoint.Duration,
		TimeHumanReadable:     formatDuration(algorithmRes.BestWay.StartingPoint.Duration),
	})
	for _, location := range algorithmRes.BestWay.Genes {
		route = append(route, LocationRes{
			Address:               location.Address,
			IsStarting:            false,
			DistanceToNextPoint:   location.Distance,
			DistanceHumanReadable: formatDistance(int(location.Distance)),
			TimeSeconds:           location.Duration,
			TimeHumanReadable:     formatDuration(location.Duration),
		})
	}

	return Response{
		Route:                      route,
		TotalDistance:              algorithmRes.BestWay.TotalDistance,
		TotalDistanceHumanReadable: formatDistance(int(algorithmRes.BestWay.TotalDistance)),
		TotalTimeHumanReadable:     formatDuration(algorithmRes.BestWay.TotalDuration),
		PopulationSize:             algorithmRes.PopulationSize,
		MaxGenerations:             algorithmRes.MaxGenerations,
		ElitismNumber:              algorithmRes.ElitismNumber,
		MutationRate:               algorithmRes.MutationRate,
	}
}

func formatDuration(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	minutes := int(duration.Minutes())
	remainingSeconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%dm%ds", minutes, remainingSeconds)
}

func formatDistance(meters int) string {
	if meters < 1000 {
		return fmt.Sprintf("%d m", meters)
	}
	kilometers := float64(meters) / 1000
	return fmt.Sprintf("%.1f km", kilometers)
}
