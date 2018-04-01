package jsonlogic

// LogicApplier can apply logic(rule) on data, based on defined operations
type LogicApplier interface {
	Apply(rule interface{}, data interface{}) (interface{}, error)
	AddOperation(symbol string, op Operator) error
}

// RuleType is the type of rule
type RuleType interface{}

// DataType is the type of data
type DataType interface{}

// Operator is an interface that can be added to LogicApplier
type Operator interface {
	Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error)
}
