package HeatersController

import (
	"errors"

	"github.com/stianeikeland/go-rpio"
)

var heater1, heater2, heater3 rpio.Pin

func init() {
	rpio.Open()

	heater1 := rpio.Pin(26)
	heater2 := rpio.Pin(20)
	heater3 := rpio.Pin(21)

	heater1.Output()
	heater2.Output()
	heater3.Output()

	heater1.High()
	heater2.High()
	heater3.High()
}

// SetNumberOfWorkingHeaters is used to set the number of working heaters
func SetNumberOfWorkingHeaters(num int) error {
	if num == 0 {
		heater1.High()
		heater2.High()
		heater3.High()
	} else if num == 1 {
		heater1.Low()
		heater2.High()
		heater3.High()
	} else if num == 2 {
		heater1.Low()
		heater2.Low()
		heater3.High()
	} else if num == 3 {
		heater1.Low()
		heater2.Low()
		heater3.Low()
	} else {
		return errors.New("The num argument should be between 0 and 3 inclusive")
	}

	return nil
}
