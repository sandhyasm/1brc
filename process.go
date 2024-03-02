package brc

import (
	"bufio"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"strconv"
	"strings"
)

// Paris=1.2/20.6/12.2

type temperateData struct {
	temperatureMin float64
	temperatureMax float64
	temperatureAvg float64
	cityCount      float64
}

func Process(r io.Reader) (string, error) {
	var cityData = make(map[string]*temperateData)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	wg := errgroup.Group{}
	for scanner.Scan() {
		wg.Go(func() error {
			line := scanner.Text()
			splitData := strings.Split(line, ";")
			city := splitData[0]
			temperature, err := strconv.ParseFloat(splitData[1], 64)
			if err != nil {
				return fmt.Errorf("Error in parsing temperature")
			}

			if cityData[city] != nil {
				if cityData[city].temperatureMin > float64(temperature) {
					cityData[city].temperatureMax = cityData[city].temperatureMin
					cityData[city].temperatureMin = float64(temperature)
				} else if cityData[city].temperatureMax < float64(temperature) {
					cityData[city].temperatureMin = cityData[city].temperatureMax
					cityData[city].temperatureMax = float64(temperature)
				}
				cityData[city].cityCount = cityData[city].cityCount + 1
				cityData[city].temperatureAvg = cityData[city].temperatureAvg + float64(temperature)
			} else {
				cityData[city] = &temperateData{
					temperatureMin: float64(temperature),
					temperatureAvg: float64(temperature),
					cityCount:      1,
				}
			}
			return err
		})
	}

	err := wg.Wait()
	if err != nil {
		return "", err
	}

	var avgData string
	for k, object := range cityData {
		avgData += fmt.Sprintf("%s=%f/%f/%f\n", k, object.temperatureMin, object.temperatureMax, object.temperatureAvg/object.cityCount)
	}

	return avgData, nil
}
