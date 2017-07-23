package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// TryReadTemperature is used to access the current temperature from the temperature sensor
func TryReadTemperature() (float32, error) {
	var error error
	var bytes []byte
	bytes, error = ioutil.ReadFile("/sys/bus/w1/devices/28-8000001fa053/w1_slave")

	if error != nil {
		return 0, error
	}

	var wholeText = string(bytes)
	var lines = strings.Split(wholeText, "\n")
	if len(lines) != 2 {
		return 0, error
	}

	if strings.HasSuffix(lines[0], "NO") {
		return 0, error
	}

	var indexOfT = strings.Index(lines[1], "t=")
	if indexOfT == -1 {
		return 0, error
	}

	var temperatureAsString = string(lines[1][indexOfT+2:])
	var temperature float64
	temperature, error = strconv.ParseFloat(temperatureAsString, 32)
	if error != nil {
		return 0, error
	}

	return float32(temperature / 1000), nil
}
