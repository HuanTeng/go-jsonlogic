package jsonlogic

import (
	"log"
)

type jsonLogic struct {
	ops map[string]Operator
}

func (jl *jsonLogic) Apply(rule interface{}, data interface{}) (interface{}, error) {
	return jl.apply(rule, data)
}

func (jl *jsonLogic) AddOperation(symbol string, op Operator) error {
	if op == nil {
		delete(jl.ops, symbol)
	} else {
		jl.ops[symbol] = op
	}
	return nil
}

func (jl *jsonLogic) mustAddOperation(symbol string, op Operator) {
	if err := jl.AddOperation(symbol, op); err != nil {
		log.Fatal(err)
	}
	return
}

// NewEmptyLogic makes an empty LogicApplier without operators
func NewEmptyLogic() LogicApplier {
	return &jsonLogic{ops: make(map[string]Operator)}
}

// NewJSONLogic makes a new LogicApplier with default jsonlogic operators
func NewJSONLogic() LogicApplier {
	jl := &jsonLogic{ops: make(map[string]Operator)}
	jl.mustAddOperation("+", opAdd{})
	jl.mustAddOperation("-", opSub{})
	jl.mustAddOperation("*", opMul{})
	jl.mustAddOperation("/", opDiv{})
	jl.mustAddOperation("%", opMod{})
	jl.mustAddOperation(">", opGreater{})
	jl.mustAddOperation(">=", opGreaterEqual{})
	jl.mustAddOperation("<", opLess{})
	jl.mustAddOperation("<=", opLessEqual{})
	jl.mustAddOperation("min", opMin{})
	jl.mustAddOperation("max", opMax{})
	jl.mustAddOperation("var", opVar{})
	jl.mustAddOperation("if", opIf{})
	jl.mustAddOperation("===", opStrictEqual{})
	jl.mustAddOperation("!==", opStrictNEqual{})
	jl.mustAddOperation("!", opNot{})
	jl.mustAddOperation("!!", opNotNot{})
	jl.mustAddOperation("or", opOr{})
	jl.mustAddOperation("and", opAnd{})

	jl.mustAddOperation("_default_aggregator", opAndBool{})
	jl.mustAddOperation("_quick_access", opVarStrictEqual{})

	return jl
}
