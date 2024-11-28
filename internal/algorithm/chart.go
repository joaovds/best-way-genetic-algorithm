package algorithm

import (
	"bytes"
	"html/template"
	"strings"
)

type SeriesData struct {
	Value float64 `json:"value"`
}

func (a *Algorithm) ChartHTML() (string, error) {
	fitnessBetterDataSerie := make([]SeriesData, 0, len(a.stats))
	distanceBetterDataSerie := make([]SeriesData, 0, len(a.stats))
	fitnessMiddleDataSerie := make([]SeriesData, 0, len(a.stats))
	distanceMiddleDataSerie := make([]SeriesData, 0, len(a.stats))
	fitnessWorseDataSerie := make([]SeriesData, 0, len(a.stats))
	distanceWorseDataSerie := make([]SeriesData, 0, len(a.stats))

	for _, stats := range a.stats {
		fitnessBetterDataSerie = append(fitnessBetterDataSerie, SeriesData{Value: stats.better.fitness})
		fitnessMiddleDataSerie = append(fitnessMiddleDataSerie, SeriesData{Value: stats.middle.fitness})
		fitnessWorseDataSerie = append(fitnessWorseDataSerie, SeriesData{Value: stats.worse.fitness})
		distanceBetterDataSerie = append(distanceBetterDataSerie, SeriesData{Value: stats.better.distance})
		distanceMiddleDataSerie = append(distanceMiddleDataSerie, SeriesData{Value: stats.middle.distance})
		distanceWorseDataSerie = append(distanceWorseDataSerie, SeriesData{Value: stats.worse.distance})
	}

	xAxisData := make([]int, a.config.MaxGenerations)
	for i := range a.config.MaxGenerations {
		xAxisData[i] = i + 1
	}

	fitnessOptions := map[string]any{
		"legend": map[string]any{},
		"title":  map[string]any{"text": "Convergence - Fitness"},
		"yAxis":  []any{struct{}{}},
		"xAxis": []map[string][]int{
			{
				"data": xAxisData,
			},
		},
		"tooltip": map[string]any{
			"show":    true,
			"trigger": "axis",
		},
		"series": []map[string]any{
			{
				"name":       "Better",
				"type":       "line",
				"smooth":     true,
				"showSymbol": true,
				"symbol":     "diamond",
				"symbolSize": 10,
				"data":       fitnessBetterDataSerie,
			},
			{
				"name":       "Middle",
				"type":       "line",
				"smooth":     true,
				"showSymbol": true,
				"symbol":     "diamond",
				"symbolSize": 10,
				"data":       fitnessMiddleDataSerie,
			},
			{
				"name":       "Worse",
				"type":       "line",
				"smooth":     true,
				"showSymbol": true,
				"symbol":     "diamond",
				"symbolSize": 10,
				"data":       fitnessWorseDataSerie,
			},
		},
	}

	distanceOptions := map[string]any{
		"legend": map[string]any{},
		"title":  map[string]any{"text": "Convergence - Distance"},
		"yAxis":  []any{struct{}{}},
		"xAxis": []map[string][]int{
			{
				"data": xAxisData,
			},
		},
		"tooltip": map[string]any{
			"show":    true,
			"trigger": "axis",
		},
		"series": []map[string]any{
			{
				"name":       "Better",
				"type":       "line",
				"smooth":     true,
				"showSymbol": true,
				"symbol":     "diamond",
				"symbolSize": 10,
				"data":       distanceBetterDataSerie,
			},
			{
				"name":       "Middle",
				"type":       "line",
				"smooth":     true,
				"showSymbol": true,
				"symbol":     "diamond",
				"symbolSize": 10,
				"data":       distanceMiddleDataSerie,
			},
			{
				"name":       "Worse",
				"type":       "line",
				"smooth":     true,
				"showSymbol": true,
				"symbol":     "diamond",
				"symbolSize": 10,
				"data":       distanceWorseDataSerie,
			},
		},
	}

	axisData := make([]int, a.chromosomeSize+1)
	axisData[0] = a.response.BestWay.StartingPoint.GetID()
	for i := range a.response.BestWay.Genes {
		axisData[i+1] = a.response.BestWay.Genes[i].GetID()
	}
	graphData := make([]float64, a.chromosomeSize+1)
	graphData[0] = a.response.BestWay.StartingPoint.Distance
	for i := range a.chromosomeSize {
		graphData[i+1] = a.response.BestWay.Genes[i].Distance
	}
	links := make([]map[string]any, a.chromosomeSize+1)
	for i := range a.chromosomeSize + 1 {
		if i == a.chromosomeSize {
			links[i] = map[string]any{"source": i, "target": 0}
		} else {
			links[i] = map[string]any{"source": i, "target": i + 1}
		}
	}

	distanceGraphOptions := map[string]any{
		"legend": map[string]any{},
		"title":  map[string]any{"text": "Graph - Distance"},
		"yAxis":  []map[string]any{{"type": "value"}},
		"xAxis": map[string]any{
			"type":        "category",
			"boundaryGap": false,
			"data":        axisData,
		},
		"tooltip": map[string]any{
			"show":    true,
			"trigger": "axis",
		},
		"series": []map[string]any{
			{
				"type":             "graph",
				"layout":           "none",
				"coordinateSystem": "cartesian2d",
				"symbolSize":       40,
				"label": map[string]any{
					"show": true,
				},
				"edgeSymbol":     []string{"circle", "arrow"},
				"edgeSymbolSize": []int{4, 10},
				"data":           graphData,
				"links":          links,
			},
		},
	}

	return renderTemplate(fitnessOptions, distanceOptions, distanceGraphOptions)
}

func renderTemplate(fitnessOptions, distanceOptions, distanceGraphOptions any) (string, error) {
	tmpl, err := template.New("chart").Parse(`
		<!-- <script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"></script> -->
			<div class="item" id="fitnessChart" style="width:900px;height:500px;"></div>
			<div class="item" id="distanceChart" style="width:900px;height:500px;"></div>
			<div class="item" id="distanceGraph" style="width:900px;height:500px;"></div>
		<script type="text/javascript">
		"use strict";
			(function() {
				let fitnessChart = echarts.init(document.getElementById('fitnessChart'), "dark", { renderer: "canvas" });
				let distanceChart = echarts.init(document.getElementById('distanceChart'), "dark", { renderer: "canvas" });
				let distanceGraph = echarts.init(document.getElementById('distanceGraph'), "dark", { renderer: "canvas" });

				let fitnessOption = {{.FitnessOption}};
				let distanceOption = {{.DistanceOption}};
				let distanceGraphOption = {{.DistanceGraphOption}};

				fitnessChart.setOption(fitnessOption);
				distanceChart.setOption(distanceOption);
				distanceGraph.setOption(distanceGraphOption);
			})();
		</script>
		<style>
		.container {margin-top:30px; display: flex;justify-content: center;align-items: center;}
		.item {margin: auto;}
		</style>
	`)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	data := map[string]any{
		"FitnessOption":       fitnessOptions,
		"DistanceOption":      distanceOptions,
		"DistanceGraphOption": distanceGraphOptions,
	}
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	cleanHTML := strings.Join(strings.Fields(tpl.String()), " ")
	return cleanHTML, nil
}
