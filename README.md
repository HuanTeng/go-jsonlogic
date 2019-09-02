# go-jsonlogic

[![Travis CI](https://travis-ci.org/HuanTeng/go-jsonlogic.svg?branch=master)](https://travis-ci.org/HuanTeng/go-jsonlogic)
[![Go Report Card](https://goreportcard.com/badge/github.com/HuanTeng/go-jsonlogic)](https://goreportcard.com/report/github.com/HuanTeng/go-jsonlogic)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/3e9df51b227c47b6b903a2a78ae62072)](https://www.codacy.com/app/the729/go-jsonlogic?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=HuanTeng/go-jsonlogic&amp;utm_campaign=Badge_Grade)

Golang implementation of JsonLogic (jsonlogic.com), which is an abstract syntax tree (AST) represented as a JSON object. 

Custom operators are supported.

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
	p0, ok0 := params[0].(float64)
	p1, ok1 := params[1].(float64)
	if !ok0 || !ok1 {
		// here we ignore the case where parameters are strings like "15"
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

	// numbers output should always be float64
	return float64(out), nil
}

// Create an instance of `jsonlogic`, and register gcd operator
jl := jsonlogic.NewJSONLogic()
jl.AddOperation("gcd", gcd{})

// Use `gcd` as a normal operator
var rule interface{}
json.Unmarshal([]byte(`{"gcd": [15, 25]}`), &rule)

got, err := jl.Apply(rule, nil)
if err != nil {
	// handle error
}

fmt.Println(got)
// Output: 5
```
