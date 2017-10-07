package tiedot

import (
	"thermo-smart/config"
	"thermo-smart/data/models"
	"time"
)

func NewTemperaturesRepository(c *Context) *TemperaturesRepository {
	return &TemperaturesRepository{
		context: c,
	}
}

type TemperaturesRepository struct {
	context *Context
}

func (self *TemperaturesRepository) Find(query *Query) []models.Temperature {
	documents := query.result

	result := make([]models.Temperature, 0)
	// To use the document itself, simply read it back
	for _, document := range documents {
		t, _ := time.Parse(config.Time.RFC3339, document[models.TemperaturesSchema.Date].(string))
		result = append(result, models.Temperature{
			Indication: document[models.TemperaturesSchema.Indication].(float64),
			Date:       t,
			Threshold:  document[models.TemperaturesSchema.Threshold].(float64),
		})
	}

	return result
}

func (self *TemperaturesRepository) Insert(item *models.Temperature) error {
	document := make(map[string]interface{})
	document[models.TemperaturesSchema.Indication] = item.Indication
	document[models.TemperaturesSchema.Date] = item.Date
	document[models.TemperaturesSchema.Threshold] = item.Threshold

	self.context.Temperatures().Insert(document)

	return nil
}

func (self *TemperaturesRepository) InitQuery() *Query {
	return NewQuery(self.context.Temperatures())
}
