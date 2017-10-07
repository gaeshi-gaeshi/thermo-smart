package HeatersController

import (
	"errors"

	"github.com/stianeikeland/go-rpio"
)

var heaters [3]rpio.Pin

func init() {
	rpio.Open()

	heaters[0] = rpio.Pin(26)
	heaters[1] = rpio.Pin(20)
	heaters[2] = rpio.Pin(21)

	heaters[0].Output()
	heaters[1].Output()
	heaters[2].Output()

	heaters[0].High()
	heaters[1].High()
	heaters[2].High()
}

// SetNumberOfWorkingHeaters is used to set the number of working heaters
func SetNumberOfWorkingHeaters(num int) error {
	if num == 0 {
		heaters[0].High()
		heaters[1].High()
		heaters[2].High()
	} else if num == 1 {
		heaters[0].Low()
		heaters[1].High()
		heaters[2].High()
	} else if num == 2 {
		heaters[0].Low()
		heaters[1].Low()
		heaters[2].High()
	} else if num == 3 {
		heaters[0].Low()
		heaters[1].Low()
		heaters[2].Low()
	} else {
		return errors.New("The num argument should be between 0 and 3 inclusive")
	}

	return nil
}

// TurnOn is used to turn on single heater
func TurnOn(heaterID int) {
	heaters[heaterID-1].Low()
}

// TurnOff is used to turn off single heater
func TurnOff(heaterID int) {
	heaters[heaterID-1].High()
}

// GetNumberOfWorkingHeaters is used to get the number of heaters currently working
func GetNumberOfWorkingHeaters() int {
	result := 0
	if heaters[0].Read() == rpio.Low {
		result++
	}

	if heaters[1].Read() == rpio.Low {
		result++
	}

	if heaters[2].Read() == rpio.Low {
		result++
	}

	return result
}
