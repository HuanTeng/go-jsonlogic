package jsonlogic_test

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/HuanTeng/go-jsonlogic"
)

func ExampleApply() {
	var rule, data interface{}

	json.Unmarshal([]byte(`
{"if" : [
	{"<": [{"var":"temp"}, 0] }, "freezing",
	{"<": [{"var":"temp"}, 100] }, "liquid",
	"gas"
]}`), &rule)

	json.Unmarshal([]byte(`{"temp":55}`), &data)

	got, err := jsonlogic.Apply(rule, data)
	if err != nil {
		// handle error
	}

	fmt.Println(got)
	// Output: liquid
}

type gcd struct{}

func (gcd) Operate(applier jsonlogic.LogicApplier, data jsonlogic.DataType, params []jsonlogic.RuleType) (jsonlogic.DataType, error) {
	if len(params) != 2 {
		return nil, errors.New("only support 2 params")
	}
	p0, ok0 := params[0].(float64)
	p1, ok1 := params[1].(float64)
	if !ok0 || !ok1 {
		return nil, errors.New("params should be numbers")
	}

	var gcdFunc func(a int, b int) int
	gcdFunc = func(a int, b int) int {
		if b == 0 {
			return a
		}
		return gcdFunc(b, a%b)
	}
	out := gcdFunc(int(p0), int(p1))

	return float64(out), nil
}

func ExampleCustomOperator() {
	jl := jsonlogic.NewJSONLogic()
	jl.AddOperation("gcd", gcd{})

	var rule interface{}
	json.Unmarshal([]byte(`{"gcd": [15, 25]}`), &rule)

	got, err := jl.Apply(rule, nil)
	if err != nil {
		// handle error
	}

	fmt.Println(got)
	// Output: 5
}
