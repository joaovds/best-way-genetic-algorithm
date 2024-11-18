package algorithm

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/joaovds/best-way-genetic-algorithm/internal/distance"
	"github.com/joaovds/best-way-genetic-algorithm/internal/operation"
)

type (
	Algorithm struct {
		config         *Config
		startingPoint  *core.Location
		locations      []*core.Location
		stats          []generationStats
		populationSize int
		chromosomeSize int
		maxGenerations int
	}

	generationStats struct {
		betterFitness, middleFitness, worseFitness float64
	}
)

func NewAlgorithm(config *Config, startingPoint *core.Location, locations []*core.Location) *Algorithm {
	chromosomeSize := len(locations)
	populationSize := 1

	if chromosomeSize <= 5 {
		for i := range chromosomeSize {
			populationSize = populationSize * (i + 1)
		}
	} else {
		populationSize = chromosomeSize * 100
	}

	if populationSize > config.MaxPopulationSize {
		populationSize = config.MaxPopulationSize
	}

	return &Algorithm{
		config:         config,
		startingPoint:  startingPoint,
		locations:      locations,
		populationSize: populationSize,
		chromosomeSize: chromosomeSize,
	}
}

func (a *Algorithm) Run() {
	distanceCalculator := distance.NewInBatchCalculator()
	distanceCalculator.CalculateDistances(a.locations, core.GetCacheInstance())
	selection := operation.NewRouletteWheelSelection()
	crossover := operation.NewPMX()
	mutation := operation.NewMutation()

	population := core.GenerateInitialPopulation(a.populationSize, a.startingPoint, a.locations, core.GetCacheInstance)

	for range a.config.MaxGenerations {
		population.EvaluateFitness()
		population.SortByFitness()

		a.stats = append(a.stats, generationStats{
			betterFitness: population.Chromosomes[0].Fitness,
			middleFitness: population.Chromosomes[population.GetSize()/2].Fitness,
			worseFitness:  population.Chromosomes[population.GetSize()-1].Fitness,
		})

		population = population.GenerateNextGeration(selection, crossover, mutation)
	}
}

func (a *Algorithm) RenderChart() {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Convergence"}), charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}))
	generations := make([]int, len(a.stats))
	betterData := make([]opts.LineData, len(a.stats))
	middleData := make([]opts.LineData, len(a.stats))
	worseData := make([]opts.LineData, len(a.stats))
	for i := range len(a.stats) {
		generations[i] = i + 1
		betterData[i] = opts.LineData{Value: a.stats[i].betterFitness}
		middleData[i] = opts.LineData{Value: a.stats[i].middleFitness}
		worseData[i] = opts.LineData{Value: a.stats[i].worseFitness}
	}
	line.SetXAxis(generations).
		AddSeries("Better", betterData).
		AddSeries("Middle", middleData).
		AddSeries("Worse", worseData).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 10, Symbol: "diamond"},
		), charts.WithAreaStyleOpts(opts.AreaStyle{
			Opacity: 0.1,
		}))

	f, err := os.Create("convergence_graph.html")
	if err != nil {
		panic(err)
	}
	line.Render(f)
}
