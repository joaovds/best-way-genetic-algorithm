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

	return renderTemplate(fitnessOptions, distanceOptions)
}

func renderTemplate(fitnessOptions, distanceOptions any) (string, error) {
	tmpl, err := template.New("chart").Parse(`
		<!-- <script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"></script> -->
			<div class="item" id="fitnessChart" style="width:900px;height:500px;"></div>
			<div class="item" id="distanceChart" style="width:900px;height:500px;"></div>
		<script type="text/javascript">
		"use strict";
			(function() {
				let fitnessChart = echarts.init(document.getElementById('fitnessChart'), "dark", { renderer: "canvas" });
				let distanceChart = echarts.init(document.getElementById('distanceChart'), "dark", { renderer: "canvas" });

				let fitnessOption = {{.FitnessOption}};
				let distanceOption = {{.DistanceOption}};

				fitnessChart.setOption(fitnessOption);
				distanceChart.setOption(distanceOption);
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
		"FitnessOption":  fitnessOptions,
		"DistanceOption": distanceOptions,
	}
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	cleanHTML := strings.Join(strings.Fields(tpl.String()), " ")
	return cleanHTML, nil
}
