package mongodb

import "github.com/gaeshi-gaeshi/thermo-smart/data/models"

func NewTemperaturesRepository(c *Context) *TemperaturesRepository {
	return &TemperaturesRepository{
		context: c,
	}
}

type TemperaturesRepository struct {
	context *Context
}

func (self *TemperaturesRepository) Find(query *Query) []models.Temperature {
	var result []models.Temperature
	query.mongoQuery.All(&result)

	return result
}

func (self *TemperaturesRepository) Insert(item *models.Temperature) error {
	err := self.context.Temperatures().Insert(item)
	return err
}

func (self *TemperaturesRepository) InitQuery() *Query {
	return NewQuery(self.context.Temperatures())
}
