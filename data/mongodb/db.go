package mongodb

import (
	"thermo-smart/config"
	"thermo-smart/data/models"

	"gopkg.in/mgo.v2"
)

type Context struct {
	session *mgo.Session
}

func (self *Context) Open() error {
	session, err := mgo.Dial("127.0.0.1")
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
	return self.session.DB(config.App.DbName).C(models.TemperaturesSchema.Name)
}
