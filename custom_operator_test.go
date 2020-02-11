package jsonlogic_test

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/HuanTeng/go-jsonlogic"
)

// gcd operator
type gcd struct{}

// Operate implements `jsonlogic.Operator`
func (gcd) Operate(applier jsonlogic.LogicApplier, data jsonlogic.DataType, params []jsonlogic.RuleType) (jsonlogic.DataType, error) {
	if len(params) != 2 {
		return nil, errors.New("only support 2 params")
	}

	var (
		p0, p1 interface{}
		err    error
	)
	// apply jsonlogic to each parameters recursively
	if p0, err = applier.Apply(params[0], data); err != nil {
		return nil, err
	}
	if p1, err = applier.Apply(params[1], data); err != nil {
		return nil, err
	}
	p0f, ok0 := p0.(float64)
	p1f, ok1 := p1.(float64)
	if !ok0 || !ok1 {
		return nil, errors.New("params should be numbers")
	}

	// recursive GCD function
	var gcdFunc func(a int, b int) int
	gcdFunc = func(a int, b int) int {
		if b == 0 {
			return a
		}
		return gcdFunc(b, a%b)
	}
	out := gcdFunc(int(p0f), int(p1f))

	// to output a number, always use float64
	return float64(out), nil
}

func ExampleOperator() {
	jl := jsonlogic.NewJSONLogic()
	jl.AddOperation("gcd", gcd{})

	var rule interface{}
	json.Unmarshal([]byte(`{"gcd": [{"+": [14, 1]}, 25]}`), &rule)

	got, err := jl.Apply(rule, nil)
	if err != nil {
		// handle error
	}

	fmt.Println(got)
	// Output: 5
}
