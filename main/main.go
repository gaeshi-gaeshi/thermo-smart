package main

import (
	//"fmt"

	//"github.com/gaeshi-gaeshi/thermo-smart/TemperatureSensorController"
	"time"

	"github.com/gaeshi-gaeshi/thermo-smart/HeatersController"
)

func main() {
	HeatersController.SetNumberOfWorkingHeaters(1)
	time.Sleep(time.Second)
	HeatersController.SetNumberOfWorkingHeaters(2)
	time.Sleep(time.Second)
	HeatersController.SetNumberOfWorkingHeaters(3)
	time.Sleep(time.Second)
	HeatersController.SetNumberOfWorkingHeaters(2)
	time.Sleep(time.Second)
	HeatersController.SetNumberOfWorkingHeaters(1)
	time.Sleep(time.Second)
	HeatersController.SetNumberOfWorkingHeaters(0)
	//var temp, error = TemperatureSensorController.ReadTemperature()
	//if error != nil {
	//	fmt.Println(error)
	//	return
	//}

	//fmt.Println(temp)
}
