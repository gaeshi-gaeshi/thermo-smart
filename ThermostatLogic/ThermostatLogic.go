package ThermostatLogic

import (
	"fmt"
	"time"

	"github.com/gaeshi-gaeshi/thermo-smart/HeatersController"
)

// Simple is very simple logic used to control the heaters based on the current and target temperatures
func Simple(targetTemperature float32, getCurrentTemperature func() float32, printCurrentTemperature func(float32)) {
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