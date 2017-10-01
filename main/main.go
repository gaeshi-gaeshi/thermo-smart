package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gaeshi-gaeshi/thermo-smart/HeatersController"
	"github.com/gaeshi-gaeshi/thermo-smart/TemperatureSensorController"
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

	for {
		currentTemperature := getCurrentTemperature()

		temperatureDifference := targetTemperature - currentTemperature
		if temperatureDifference > 1 {
			HeatersController.SetNumberOfWorkingHeaters(3)
		} else if temperatureDifference <= 1 && temperatureDifference > 0 {
			HeatersController.SetNumberOfWorkingHeaters(2)
		} else if temperatureDifference <= 0 && temperatureDifference > -1 {
			HeatersController.SetNumberOfWorkingHeaters(1)
		} else {
			HeatersController.SetNumberOfWorkingHeaters(0)
		}

		printCurrentTemperature(currentTemperature)
		fmt.Printf("Currently working heaters - %d\n", HeatersController.GetNumberOfWorkingHeaters())

		time.Sleep(time.Minute)
	}
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
