package algorithm

import (
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
	"github.com/joaovds/best-way-genetic-algorithm/internal/distance"
	"github.com/joaovds/best-way-genetic-algorithm/internal/operation"
)

const (
	MAX_POPULATION_SIZE = 7000
	MAX_GENERATIONS     = 300
)

type (
	Algorithm struct {
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

func NewAlgorithm(startingPoint *core.Location, locations []*core.Location) *Algorithm {
	chromosomeSize := len(locations)
	populationSize := 1

	if chromosomeSize <= 5 {
		for i := range chromosomeSize {
			populationSize = populationSize * (i + 1)
		}
	} else {
		populationSize = chromosomeSize * 100
	}

	if populationSize > MAX_POPULATION_SIZE {
		populationSize = MAX_POPULATION_SIZE
	}

	return &Algorithm{
		startingPoint:  startingPoint,
		locations:      locations,
		populationSize: populationSize,
		chromosomeSize: chromosomeSize,
	}
}

func (a *Algorithm) Run() {
	distanceCalculator := distance.NewSimpleDistanceCalculator()
	selection := operation.NewRouletteWheelSelection()
	crossover := operation.NewPMX()
	mutation := operation.NewMutation()

	population := core.GenerateInitialPopulation(a.populationSize, a.startingPoint, a.locations, core.GetCacheInstance)

	for range MAX_GENERATIONS {
		population.EvaluateFitness(distanceCalculator)
		population.SortByFitness()

		a.stats = append(a.stats, generationStats{
			betterFitness: population.Chromosomes[0].Fitness,
			middleFitness: population.Chromosomes[population.GetSize()/2].Fitness,
			worseFitness:  population.Chromosomes[population.GetSize()-1].Fitness,
		})

		// fmt.Println("Location: ", a.startingPoint)
		// fmt.Println("Locales:")
		// for _, localion := range a.locations {
		// 	fmt.Println(localion)
		// }
		//
		// for _, c := range population.Chromosomes {
		// 	fmt.Println("\n----- ... ----- \nStart:", c.StartingPoint)
		// 	fmt.Println("\nGenes:")
		// 	for _, gene := range c.Genes {
		// 		fmt.Println(gene)
		// 	}
		// 	fmt.Println("Fitness:", c.Fitness)
		// }
		// fmt.Println("\nPopulation Size:", population.GetSize())
		// fmt.Println("\nPopulation Total Fitness:", population.TotalFitness)
		fmt.Println(population.Chromosomes[0].Fitness)
		fmt.Println(population.Chromosomes[0].SurvivalCount)

		population = population.GenerateNextGeration(selection, crossover, mutation)
	}
	fmt.Print(population.Chromosomes[0].StartingPoint.GetID(), population.Chromosomes[0].StartingPoint.Distance)
	for _, gene := range population.Chromosomes[0].Genes {
		fmt.Print("->", gene.GetID(), gene.Distance)
	}
	fmt.Println(distance.Counter)
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
