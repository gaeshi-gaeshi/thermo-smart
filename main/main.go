package main

import (
	"fmt"

	"github.com/gaeshi-gaeshi/thermo-smart/TemperatureSensorController"
)

func main() {
	var temp, error = TemperatureSensorController.ReadTemperature()
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(temp)
}
