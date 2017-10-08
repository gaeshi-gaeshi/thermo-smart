package mongodb

import (
	"github.com/gaeshi-gaeshi/thermo-smart/data/models"

	"gopkg.in/mgo.v2"
)

func NewContext(dbName string, dbUrl string) *Context {
	context := &Context{
		dbName: dbName,
		dbUrl:  dbUrl,
	}

	return context
}

type Context struct {
	session *mgo.Session
	dbName  string
	dbUrl   string
}

func (self *Context) Open() error {
	session, err := mgo.Dial(self.dbUrl)
	if err != nil {
		return err
	}

	session.SetMode(mgo.Monotonic, true)
	self.session = session

	return nil
}

func (self *Context) Close() error {
	self.session.Close()
	return nil
}

func (self *Context) Temperatures() *mgo.Collection {
	return self.session.DB(self.dbName).C(models.TemperaturesSchema.Name)
}
