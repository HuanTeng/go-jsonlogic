# go-jsonlogic

[![Travis CI](https://travis-ci.org/ShankarParimi/go-jsonlogic.svg?branch=master)](https://travis-ci.org/ShankarParimi/go-jsonlogic)
[![Go Report Card](https://goreportcard.com/badge/github.com/HuanTeng/go-jsonlogic)](https://goreportcard.com/report/github.com/ShankarParimi/go-jsonlogic)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2f0b5e62e6134373baecd36e346bdcb1)](https://www.codacy.com/manual/ShankarParimi/go-jsonlogic?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ShankarParimi/go-jsonlogic&amp;utm_campaign=Badge_Grade)
Golang implementation of JsonLogic (jsonlogic.com), which is an abstract syntax tree (AST) represented as a JSON object. 

Custom operators are supported.

Rules Validation is supported.

## Example

```golang
var rule, data interface{}

json.Unmarshal([]byte(`
{"if": [
	{"<": [{"var":"temp"}, 0] }, "freezing",
	{"<": [{"var":"temp"}, 100] }, "liquid",
	"gas"
]}
`), &rule)

json.Unmarshal([]byte(`{"temp":55}`), &data)

got, err := jsonlogic.Apply(rule, data)
if err != nil {
	// handle error
}

fmt.Println(got)
// Output: liquid
```

## Custom operators

You can add your own operator to `jsonlogic` by implementing `jsonlogic.Operator` interface and registering it to a `jsonlogic` instance.

The following example shows how to implement a greatest common divisor (gcd) operator.

```golang

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

// Create an instance of `jsonlogic`, and register gcd operator
jl := jsonlogic.NewJSONLogic()
jl.AddOperation("gcd", gcd{})

// Use `gcd` as an operator to calculate: gcd(14+1, 25)
var rule interface{}
json.Unmarshal([]byte(`{"gcd": [{"+": [14, 1]}, 25]}`), &rule)

got, err := jl.Apply(rule, nil)
if err != nil {
	// handle error
}

fmt.Println(got)
// Output: 5
```
