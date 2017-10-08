package tiedot

import (
	"github.com/gaeshi-gaeshi/thermo-smart/data/models"

	"github.com/HouzuoGuo/tiedot/db"
)

func NewContext(dbDirectory string) *Context {
	context := &Context{
		dbDirectory: dbDirectory,
	}

	context.setup()
	return context
}

type Context struct {
	dbDirectory string
	db          *db.DB
}

func (self *Context) Open() error {
	// (Create if not exist) open a database
	tiedotDb, err := db.OpenDB(self.dbDirectory)
	if err != nil {
		return err
	}

	self.db = tiedotDb
	return nil
}

func (self *Context) Close() error {
	// Gracefully close database, you should call Close on all opened databases
	// Otherwise background goroutines will prevent program shutdown
	err := self.db.Close()
	return err
}

func (self *Context) Temperatures() *db.Col {
	return self.db.Use(models.TemperaturesSchema.Name)
}

func (self *Context) setup() {
	if err := self.Open(); err != nil {
		panic(err)
	}

	self.createIfNotExist(models.TemperaturesSchema.Name)

	if err := self.Close(); err != nil {
		panic(err)
	}
}

func (self *Context) createIfNotExist(collection string) {
	contains := false
	for _, name := range self.db.AllCols() {
		if name == collection {
			contains = true
			break
		}
	}

	if !contains {
		if err := self.db.Create(collection); err != nil {
			panic(err)
		}
	}
}
