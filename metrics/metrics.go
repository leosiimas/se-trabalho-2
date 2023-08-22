package metrics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CollectHardwareData() (SensorData, error) {
	response, err := http.Get("http://localhost:8085/data.json")
	if err != nil {
		return SensorData{}, err
	}
	defer response.Body.Close()

	var data SensorData
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return SensorData{}, err
	}

	return data, nil
}

func sensorData(sensor SensorData) []string {
	var result []string

	if sensor.Text == "AMD Ryzen 7 5700G" {
		result = append(result, getHardwareInfo(sensor, "CPU")...)
		return result

	} else if sensor.Text == "Generic Memory" {
		result = append(result, getHardwareInfo(sensor, "Memory")...)
		return result

	} else {
		for _, children := range sensor.Children {
			result = append(result, sensorData(children)...)
		}
	}

	return result
}

func getHardwareInfo(sensor SensorData, hardwareName string) []string {
	var values []string

	for _, children := range sensor.Children {

		if children.Text == "Temperatures" || children.Text == "Powers" || children.Text == "Load" || children.Text == "Data" {
			for _, c := range children.Children {
				if c.Text == "CPU Package" || c.Text == "CPU Cores" || c.Text == "Memory" || c.Text == "Used Memory" {
					values = append(values, c.Value)
				}

			}
		}

	}

	return values
}

func sensorTitle(sensor SensorData) []string {
	var titles []string

	if sensor.Text == "AMD Ryzen 7 5700G" {
		titles = append(titles, getHardwareTitle(sensor, "CPU")...)
		return titles

	} else if sensor.Text == "Generic Memory" {
		titles = append(titles, getHardwareTitle(sensor, "CPU")...)
		return titles

	} else {
		for _, children := range sensor.Children {
			titles = append(titles, sensorTitle(children)...)
		}
	}
	return titles
}

func getHardwareTitle(sensor SensorData, hardwareName string) []string {
	var values []string

	for _, children := range sensor.Children {

		if children.Text == "Temperatures" || children.Text == "Powers" || children.Text == "Load" || children.Text == "Data" {
			for _, c := range children.Children {
				if c.Text == "CPU Package" || c.Text == "CPU Cores" || c.Text == "Memory" || c.Text == "Used Memory" {
					text := hardwareName + "," + children.Text + "," + c.Text

					values = append(values, text)
				}

			}
		}

	}

	return values
}

func HardwareMetrics(done chan bool, path PathValue) {
	var storageData [][]string

	duration := 0
	docHaveTitles := false

	defer func() {
		DataToJson(storageData, path)
	}()

	for {

		select {
		case <-done:
			return

		default:
			time.Sleep(time.Second)

			data, err := CollectHardwareData()
			if err != nil {
				fmt.Println("Erro ao acessar o servidor web do Open Hardware Monitor:", err)
				continue
			}

			if !docHaveTitles {
				titles := sensorTitle(data)
				titles = append(titles, "Time")
				storageData = append(storageData, titles)
				docHaveTitles = true
			}

			sensor := sensorData(data)
			sensor = append(sensor, fmt.Sprintf("%d", duration))
			storageData = append(storageData, sensor)

			duration++
		}

	}

}
