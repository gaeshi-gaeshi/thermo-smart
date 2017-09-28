package TemperatureSensorController

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// ReadTemperature is used to access the current temperature from the temperature sensor
func ReadTemperature() (float32, error) {
	var error error
	var bytes []byte
	var path = "/sys/bus/w1/devices/28-8000001fa053/w1_slave"

	bytes, error = ioutil.ReadFile(path)
	time.Sleep(time.Second)
	bytes, error = ioutil.ReadFile(path)

	if error != nil {
		return 0, error
	}

	var wholeText = string(bytes)
	var lines = strings.Split(wholeText, "\n")
	if strings.HasSuffix(lines[0], "NO") {
		return 0, errors.New("Read temperature from sensor failed - the sensor returned 'NO'")
	}

	var indexOfT = strings.Index(lines[1], "t=")
	if indexOfT == -1 {
		return 0, errors.New("Read temperature from sensor failed - 't=' not found")
	}

	var temperatureAsString = string(lines[1][indexOfT+2:])
	var temperature float64
	temperature, error = strconv.ParseFloat(temperatureAsString, 32)
	if error != nil {
		return 0, error
	}

	return round(float32(temperature/1000), 1)
}

func round(number float32, decimalDigitToRoundTo int) (float32, error) {
	var roundedNumberAsString = fmt.Sprintf("%."+strconv.Itoa(decimalDigitToRoundTo)+"f", number)
	var roundedNumberAsFloat64, error = strconv.ParseFloat(roundedNumberAsString, 32)
	if error != nil {
		return 0, error
	}

	return float32(roundedNumberAsFloat64), nil
}
