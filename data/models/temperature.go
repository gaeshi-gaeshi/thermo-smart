package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

var TemperaturesSchema = &schema{
	Name:       "Temperatures",
	Indication: "indication",
	Date:       "date",
	Threshold:  "threshold",
}

type Temperature struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
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
