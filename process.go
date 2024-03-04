package brc

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Paris=1.2/20.6/12.2

type temperatureData struct {
	temperatureMin float64
	temperatureMax float64
	temperatureSum float64
	cityCount      float64
}

func Process(r io.Reader) (string, error) {
	var cityData = make(map[string]*temperatureData)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		splitData := strings.Split(line, ";")
		city := splitData[0]
		temperature, _ := strconv.ParseFloat(splitData[1], 64)

		if data, ok := cityData[city]; ok {
			if data.temperatureMin > temperature {
				data.temperatureMin = temperature
			}
			if data.temperatureMax < temperature {
				data.temperatureMax = temperature
			}
			data.temperatureSum += temperature
			data.cityCount++
		} else {
			cityData[city] = &temperatureData{
				temperatureMin: temperature,
				temperatureMax: temperature,
				temperatureSum: temperature,
				cityCount:      1,
			}
		}
	}

	var avgData string
	for k, data := range cityData {
		avgData += fmt.Sprintf("%s=%f/%f/%f\n", k, data.temperatureMin, data.temperatureMax, data.temperatureSum/data.cityCount)
	}

	return avgData, nil
}
