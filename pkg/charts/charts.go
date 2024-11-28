package charts

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"
)

type SeriesData struct {
	Value float64 `json:"value"`
}

func ChartHTML(title string, x []int, y []float64) (string, error) {
	dataSerie := make([]SeriesData, len(y))
	for i := range y {
		dataSerie[i] = SeriesData{Value: y[i]}
	}

	options := map[string]interface{}{
		"legend": map[string]interface{}{},
		"title":  map[string]interface{}{"text": title},
		"yAxis":  []interface{}{map[string]interface{}{}},
		"xAxis": []map[string]interface{}{
			{
				"data": x,
			},
		},
		"tooltip": map[string]interface{}{
			"show":    true,
			"trigger": "axis",
		},
		"series": []map[string]interface{}{
			{
				"name":       "",
				"type":       "line",
				"smooth":     true,
				"showSymbol": true,
				"symbol":     "diamond",
				"symbolSize": 10,
				"data":       dataSerie,
			},
		},
	}

	return render(options)
}

func render(options map[string]any) (string, error) {
	optionsJSON, err := json.Marshal(options)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("chart").Parse(`
		<!-- <script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"></script> -->
			<div class="item" id="chart" style="width:900px;height:500px;"></div>
		<script type="text/javascript">
		"use strict";
			(function() {
				let chart = echarts.init(document.getElementById('chart'), "dark", { renderer: "canvas" });

				let options = {{.Options}};

				chart.setOption(options);
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
		"Options": string(optionsJSON),
	}
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	cleanHTML := strings.Join(strings.Fields(tpl.String()), " ")
	return cleanHTML, nil
}
