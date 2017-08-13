package mongodb

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewQuery(c *mgo.Collection) *Query {
	return &Query{
		collection: c,
		query:      make(map[string]interface{}),
	}
}

type Query struct {
	field      string
	collection *mgo.Collection
	query      map[string]interface{}
	MongoQuery *Query
}

func (self *Query) Field(name string) *Query {
	self.field = name
	return self
}

func (self *Query) Equals(value interface{}) *Query {
	self.query = bson.M{self.field: value}
	return self
}

func (self *Query) Between(start interface{}, end interface{}) *Query {
	self.query = bson.M{self.field: bson.M{
		"$gt": start,
		"$lt": end,
	}}

	return self
}

func (self *Query) Build() map[string]interface{} {
	self.collection.Find(self.query).Sort("-timestamp")
	return nil
}

func (self *Query) Clear() {
	self.query = nil
	self.query = make(map[string]interface{})
}
