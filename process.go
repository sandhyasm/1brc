package brc

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// Paris=1.2/20.6/12.2

var temperatureMinData = make(map[string]float64)
var temperatureMaxData = make(map[string]float64)
var temperatureAvgData = make(map[string]float64)
var cityCountData = make(map[string]int)

func Process(r io.Reader) (string, error) {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var cityArray []string
	for scanner.Scan() {
		line := scanner.Text()
		splitData := strings.Split(line, ";")
		city := splitData[0]
		temperature, err := strconv.ParseFloat(splitData[1], 64)
		if err != nil {
			return "", err
		}
		if cityCountData[city] > 0 {
			if temperatureMinData[city] > float64(temperature) {
				temperatureMaxData[city] = temperatureMinData[city]
				temperatureMinData[city] = float64(temperature)
			} else if temperatureMaxData[city] < float64(temperature) {
				temperatureMinData[city] = temperatureMaxData[city]
				temperatureMaxData[city] = float64(temperature)
			}
		} else {
			cityArray = append(cityArray, city)
			temperatureMinData[city] = float64(temperature)
		}

		cityCountData[city] = cityCountData[city] + 1
		temperatureAvgData[city] = temperatureAvgData[city] + float64(temperature)
	}
	sort.Strings(cityArray)
	var avgData string
	for _, city := range cityArray {
		temperatureAvgData[city] = temperatureAvgData[city] / float64(cityCountData[city])

		avgData += fmt.Sprintf("%s=%f/%f/%f\n", city, temperatureMinData[city], temperatureMaxData[city], temperatureAvgData[city])
	}

	return avgData, nil
}
