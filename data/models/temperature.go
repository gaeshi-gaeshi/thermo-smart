package models

import (
	"time"
)

var TemperaturesSchema = &schema{
	Name:       "Temperatures",
	Indication: "Indication",
	Date:       "Date",
	Threshold:  "Threshold",
}

type Temperature struct {
	Indication float64
	Date       time.Time
	Threshold  float64
}

type schema struct {
	Name       string
	Indication string
	Date       string
	Threshold  string
}
