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

// Complex uses complex logic to control the temperature. It keeps 1 heater for longer period of time.
// Also Ton != Toff for the 3 heaters.
func Complex(targetTemperature float32, getCurrentTemperature func() float32, printCurrentTemperature func(float32)) {
	for {
		currentTemperature := getCurrentTemperature()

		temperatureDifference := currentTemperature - targetTemperature
		if shouldIncreaseTargetTemperature() {
			temperatureDifference--
			fmt.Println("Target temperature increased by 1 degree.")
		}

		if temperatureDifference <= 1 {
			HeatersController.TurnOn(1)
		} else if temperatureDifference >= 2 {
			HeatersController.TurnOff(1)
		}

		if temperatureDifference <= -0.5 {
			HeatersController.TurnOn(2)
		} else if temperatureDifference >= 0.5 {
			HeatersController.TurnOff(2)
		}

		if temperatureDifference <= -1 {
			HeatersController.TurnOn(3)
		} else if temperatureDifference >= 0.5 {
			HeatersController.TurnOff(3)
		}

		printCurrentTemperature(currentTemperature)
		fmt.Printf("Currently working heaters - %d\n", HeatersController.GetNumberOfWorkingHeaters())

		time.Sleep(time.Minute * 5)
	}
}

func shouldIncreaseTargetTemperature() bool {
	now := time.Now()
	if now.Hour() == 23 && now.Minute() >= 30 {
		return true
	} else if now.Hour() < 6 {
		return true
	} else if now.Hour() == 6 && now.Minute() <= 30 {
		return true
	}

	return false
}
