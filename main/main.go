package main

import (
	"fmt"
	"thermo-smart/config"
	"thermo-smart/data/models"
	"thermo-smart/data/mongodb"
	"time"
)

func main() {
	context := mongodb.NewContext(config.App.DbName, "127.0.0.1")
	unit, repository := mongodb.NewUnit(context), mongodb.NewTemperaturesRepository(context)

	unit.Begin()
	repository.Insert(&models.Temperature{Indication: 32.1, Threshold: 33.0, Date: time.Now()})
	temperatures := repository.Find(repository.InitQuery().Field(models.TemperaturesSchema.Indication).Equals(32.1).Build())
	unit.Commit()

	fmt.Printf("Temperatures:\n%v", temperatures)
}
