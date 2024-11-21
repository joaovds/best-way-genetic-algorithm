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
		Locations      []Location `json:"locations"`
		MaxPopulation  int        `json:"max_population"`
		MaxGenerations int        `json:"max_generations"`
		Elitism        int        `json:"elitism"`
		MutationRate   float64    `json:"mutation_rate"`
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
		ChartsHtml                 string        `json:"charts_html"`
		TotalDistanceHumanReadable string        `json:"total_distance_human_readable"`
		TotalTimeHumanReadable     string        `json:"total_time_human_readable"`
		AlgoritmTime               string        `json:"algorithm_time"`
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
	if l.MaxGenerations == 0 {
		l.MaxGenerations = 500
	}
	if l.MaxPopulation == 0 {
		l.MaxPopulation = 7000
	}
	if l.MutationRate == 0 {
		l.MutationRate = 0.2
	}
	if l.Elitism == 0 {
		l.Elitism = 4
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

func AlgorithmResponseToApiResponse(algorithmRes *algorithm.AlgorithmResponse, chartsHtml string, startTime time.Time) Response {
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
		AlgoritmTime:               fmt.Sprintf("%ds %03dms", int(time.Since(startTime).Seconds()), int(time.Since(startTime).Milliseconds())%1000),
		ChartsHtml:                 chartsHtml,
		TotalDistance:              algorithmRes.BestWay.TotalDistance,
		TotalDistanceHumanReadable: formatDistance(int(algorithmRes.BestWay.TotalDistance)),
		TotalTime:                  algorithmRes.BestWay.TotalDuration,
		TotalTimeHumanReadable:     formatDuration(algorithmRes.BestWay.TotalDuration),
		PopulationSize:             algorithmRes.PopulationSize,
		MaxGenerations:             algorithmRes.MaxGenerations,
		ElitismNumber:              algorithmRes.ElitismNumber,
		MutationRate:               algorithmRes.MutationRate,
	}
}

func formatDuration(seconds int) string {
	duration := time.Duration(seconds) * time.Second

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	remainingSeconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%d:%d:%d", hours, minutes, remainingSeconds)
}

func formatDistance(meters int) string {
	if meters < 1000 {
		return fmt.Sprintf("%d m", meters)
	}
	kilometers := float64(meters) / 1000
	return fmt.Sprintf("%.1f km", kilometers)
}
