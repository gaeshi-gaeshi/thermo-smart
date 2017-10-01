package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gaeshi-gaeshi/thermo-smart/HeatersController"
	"github.com/gaeshi-gaeshi/thermo-smart/TemperatureSensorController"
	"github.com/gaeshi-gaeshi/thermo-smart/ThermostatLogic"
)

func main() {
	if os.Args[1] == "--reset" {
		HeatersController.SetNumberOfWorkingHeaters(0)

		return
	}

	if os.Args[1] == "--temp" {
		currentTemperature := getCurrentTemperature()
		printCurrentTemperature(currentTemperature)

		return
	}

	targetTemperatureAsFloat64, error := strconv.ParseFloat(os.Args[1], 32)
	if error != nil {
		fmt.Println(error)
		return
	}

	targetTemperature := float32(targetTemperatureAsFloat64)

	fmt.Printf("Target temperature - %.2f\n", targetTemperature)

	ThermostatLogic.Simple(targetTemperature, getCurrentTemperature, printCurrentTemperature)
}

func getCurrentTemperature() float32 {
	currentTemperature, error := TemperatureSensorController.ReadTemperature()
	if error != nil {
		log.Fatal(error)
	}

	return currentTemperature
}

func printCurrentTemperature(currentTemperature float32) {

	fmt.Printf("Current temperature - %.2f\n", currentTemperature)
}
