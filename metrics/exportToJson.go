package metrics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/go-echarts/go-echarts/v2/opts"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func DataToJson(values [][]string, path PathValue) {

	var titles []string
	var datas []DataEntry

	column := len(values[0])
	row := len(values)

	for i := 0; i < column-1; i++ {
		titles = append(titles, values[0][i])
	}

	for j := 1; j < row; j++ {
		var d DataEntry

		for i := 0; i <= column-1; i++ {
			var metricInfo MetricInfoData

			if values[0][i] == "Time" {
				d.Language = path.Language
				d.Time = values[j][i]
			} else {

				value, measure := Parse(values[j][i])

				parts := strings.Split(values[0][i], ",")
				metricInfo.Hardware = parts[0]
				metricInfo.Metric = parts[1]
				metricInfo.Sensor = parts[2]
				metricInfo.Value = value
				metricInfo.Measure = measure
				d.Data = append(d.Data, metricInfo)
			}

		}
		datas = append(datas, d)
	}

	jsonData, err := json.MarshalIndent(datas, "", "\t")
	if err != nil {
		fmt.Println("Erro ao codificar JSON:", err)
		return
	}

	err = ioutil.WriteFile(path.OutputPath+path.Text+".json", jsonData, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever JSON no arquivo:", err)
		return
	}
}
