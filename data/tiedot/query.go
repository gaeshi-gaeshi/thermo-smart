package tiedot

import (
	"github.com/HouzuoGuo/tiedot/db"
)

func NewQuery(c *db.Col) *Query {
	return &Query{
		collection: c,
		query:      make(map[string]interface{}),
	}
}

type Query struct {
	collection *db.Col
	query      map[string]interface{}
	result     []map[string]interface{}
}

func (self *Query) Field(name string) *Query {
	self.query["in"] = []interface{}{name}
	return self
}

func (self *Query) Equals(value interface{}) *Query {
	self.query["eq"] = value
	return self
}

func (self *Query) Between(start int, end int) *Query {
	self.query["int-from"] = start
	self.query["int-to"] = end
	return self
}

func (self *Query) Build() *Query {
	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys
	if err := db.EvalQuery(self.query, self.collection, &queryResult); err != nil {
		panic(err)
	}

	result := make([]map[string]interface{}, 0)
	// To use the document itself, simply read it back
	for id := range queryResult {
		readBack, err := self.collection.Read(id)
		if err != nil {
			panic(err)
		}

		result = append(result, readBack)
	}

	self.result = result

	return self
}

func (self *Query) Clear() {
	self.query = make(map[string]interface{})
	self.result = nil
}
