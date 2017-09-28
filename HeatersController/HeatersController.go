package HeatersController

import "github.com/stianeikeland/go-rpio"

func Test() {
	rpio.Open()

	pin := rpio.Pin(26)

	pin.Output()
	pin.High()
}
