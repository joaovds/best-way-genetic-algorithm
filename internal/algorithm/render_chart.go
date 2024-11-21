package algorithm

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func (a *Algorithm) RenderChart() {
	lineFitness := charts.NewLine()
	lineDistance := charts.NewLine()
	lineDistance.Theme, lineFitness.Theme = "dark", "dark"

	lineFitness.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Convergence - Fitness"}), charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}))
	lineDistance.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Convergence - Distance"}), charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}))

	generations := make([]int, len(a.stats))
	betterFitness := make([]opts.LineData, len(a.stats))
	middleFitness := make([]opts.LineData, len(a.stats))
	worseFitness := make([]opts.LineData, len(a.stats))
	betterDistance := make([]opts.LineData, len(a.stats))
	middleDistance := make([]opts.LineData, len(a.stats))
	worseDistance := make([]opts.LineData, len(a.stats))

	for i := range len(a.stats) {
		generations[i] = i + 1
		betterFitness[i] = opts.LineData{Value: a.stats[i].better.fitness * 1000}
		middleFitness[i] = opts.LineData{Value: a.stats[i].middle.fitness * 1000}
		worseFitness[i] = opts.LineData{Value: a.stats[i].middle.fitness * 1000}
		betterDistance[i] = opts.LineData{Value: a.stats[i].better.distance}
		middleDistance[i] = opts.LineData{Value: a.stats[i].middle.distance}
		worseDistance[i] = opts.LineData{Value: a.stats[i].worse.distance}
	}

	lineFitness.SetXAxis(generations).
		AddSeries("Better", betterFitness).
		AddSeries("Middle", middleFitness).
		AddSeries("Worse", worseFitness).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 10, Symbol: "diamond"},
		), charts.WithAreaStyleOpts(opts.AreaStyle{
			Opacity: 0.1,
		}))
	lineDistance.SetXAxis(generations).
		AddSeries("Better", betterDistance).
		AddSeries("Middle", middleDistance).
		AddSeries("Worse", worseDistance).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 10, Symbol: "diamond"},
		), charts.WithAreaStyleOpts(opts.AreaStyle{
			Opacity: 0.1,
		}))

	var bufFitness, bufDistance bytes.Buffer
	lineFitness.Render(&bufFitness)
	lineDistance.Render(&bufDistance)

	f, err := os.Create("convergence_graph.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	finalHTML := fmt.Sprintf(pageHTML, bufFitness.String(), bufDistance.String())
	f.WriteString(finalHTML)
}

const pageHTML = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Convergence Graph</title>
		<script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"></script>
	</head>

	<body>
	%s
	%s
	</body>
</html>
`
