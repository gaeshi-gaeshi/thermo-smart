package tiedot

func NewUnit(c *Context) *Unit {
	return &Unit{
		context: c,
	}
}

type Unit struct {
	context *Context
	// Mapper  *Mapper
}

func (self *Unit) Begin() error {
	return self.context.Open()
}

func (self *Unit) Commit() error {
	return self.context.Close()
}

// TODO: Remove
// Thist is the generic data mapping using reflection. Currently the data mapping logic is in the repository.
// Must be decided if it will stay there, and this generic code to be removed.
// type Mapper struct {
// }

// func (m *Mapper) Map(obj interface{}, arr []map[string]interface{}) {
// 	for _, item := range arr {
// 		fillObject(obj, item)
// 	}
// }

// func fillObject(o interface{}, m map[string]interface{}) error {
// 	for k, v := range m {
// 		err := setField(o, k, v)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func setField(obj interface{}, name string, value interface{}) error {
// 	structValue := reflect.ValueOf(obj).Elem()
// 	structFieldValue := structValue.FieldByName(name)

// 	if !structFieldValue.IsValid() {
// 		return fmt.Errorf("No such field: %s in obj", name)
// 	}

// 	if !structFieldValue.CanSet() {
// 		return fmt.Errorf("Cannot set %s field value", name)
// 	}

// 	structFieldType := structFieldValue.Type()
// 	val := reflect.ValueOf(value)
// 	if structFieldType != val.Type() {
// 		return errors.New("Provided value type didn't match obj field type")
// 	}

// 	structFieldValue.Set(val)
// 	return nil
// }
