package jsonlogic

var defaultJL LogicApplier

func init() {
	defaultJL = NewJSONLogic()
}

// Apply rule on data using the default JSONLogic
func Apply(rule interface{}, data interface{}) (interface{}, error) {
	return defaultJL.Apply(rule, data)
}

// AddOperation to the default JSONLogic
func AddOperation(symbol string, op Operator) error {
	return defaultJL.AddOperation(symbol, op)
}

func Validate(rule interface{}) error {
	return defaultJL.Validate(rule)
}
