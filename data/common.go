package data

import (
	"thermo-smart/data/models"
)

type Unit interface {
	Begin() error
	Commit() error
}

type Query interface {
	Field(name string) *Query
	Equals(value interface{}) *Query
	Between(start interface{}, end interface{}) *Query
	Build() *Query
	Clear()
}

type TemperaturesRepository interface {
	Find(query *Query) ([]models.Temperature, error)
	Insert(item *models.Temperature) error
	InitQuery() *Query
}
