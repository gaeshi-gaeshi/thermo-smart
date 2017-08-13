package mongodb

func NewUnit(c *Context) *Unit {
	return &Unit{
		context: c,
	}
}

type Unit struct {
	context *Context
}

func (self *Unit) Begin() error {
	return self.context.Open()
}

func (self *Unit) Commit() error {
	return self.context.Close()
}
