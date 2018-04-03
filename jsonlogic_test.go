package jsonlogic_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	jsonlogic "github.com/HuanTeng/go-jsonlogic"
)

type TestCase struct {
	name   string
	rule   string
	data   string
	expect interface{}
}

var defaultAggTestCases = []TestCase{
	{
		name:   "nil case",
		rule:   `null`,
		data:   `[true]`,
		expect: nil,
	},
	{
		name:   "_default_agg true",
		rule:   `{"===" : [1, 1], "var": 0, "+": [1]}`,
		data:   `[true]`,
		expect: true,
	},
	{
		name:   "_default_agg false",
		rule:   `{"===" : [1, 1], "var": 0, "+": [1, -1]}`,
		data:   `[true]`,
		expect: false,
	},
}

var quickVarTestCases = []TestCase{
	{
		name:   "quick var match",
		rule:   `{"$name": "Jack"}`,
		data:   `{"name": "Jack"}`,
		expect: true,
	},
	{
		name:   "2x quick var match",
		rule:   `{"$name.first": "Jack", "$name.last": "Johnson"}`,
		data:   `{"name": {"first": "Jack", "last": "Johnson"}}`,
		expect: true,
	},
	{
		name:   "2x quick var mismatch",
		rule:   `{"$name.first": "Jack", "$name.last": "Johnson"}`,
		data:   `{"name": {"first": "Johnson", "last": "Johnson"}}`,
		expect: false,
	},
}

var errorCases = []TestCase{
	{
		name:   "empty op",
		rule:   `{"": 1}`,
		data:   `{}`,
		expect: fmt.Errorf("operator  not found"),
	},
	{
		name:   "not found op",
		rule:   `{"not_found": 1}`,
		data:   `{}`,
		expect: fmt.Errorf("operator not_found not found"),
	},
	{
		name:   "not aggregator",
		rule:   `{"+": 1, "-": 1}`,
		data:   `{}`,
		expect: fmt.Errorf("multiple keys found but default aggregator not defined"),
	},
	{
		name:   "not quick access",
		rule:   `{"$id": 1}`,
		data:   `{}`,
		expect: fmt.Errorf("quick access op not defined"),
	},
}

func runTestCases(cases []TestCase, t *testing.T) {
	var rule, data interface{}

	for _, c := range cases {
		if err := json.Unmarshal([]byte(c.rule), &rule); err != nil {
			t.Errorf("Case %s: rule error: %s", c.name, err)
		}
		if err := json.Unmarshal([]byte(c.data), &data); err != nil {
			t.Errorf("Case %s: data error: %s", c.name, err)
		}
		got, err := jsonlogic.Apply(rule, data)
		if err != nil {
			t.Errorf("Case %s: apply error: %s", c.name, err)
		} else if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("Case %s: expect %s got %s", c.name,
				spew.Sdump(c.expect), spew.Sdump(got),
			)
		}
	}
}

func TestErrorCases(t *testing.T) {
	var rule, data interface{}

	jl := jsonlogic.NewJSONLogic()

	jl.AddOperation("_default_aggregator", nil)
	jl.AddOperation("_quick_access", nil)

	for _, c := range errorCases {
		if err := json.Unmarshal([]byte(c.rule), &rule); err != nil {
			t.Errorf("Case %s: rule error: %s", c.name, err)
		}
		if err := json.Unmarshal([]byte(c.data), &data); err != nil {
			t.Errorf("Case %s: data error: %s", c.name, err)
		}
		_, err := jl.Apply(rule, data)
		if !reflect.DeepEqual(err, c.expect) {
			t.Errorf("Case %s: expect error %s got %s", c.name,
				spew.Sdump(c.expect), spew.Sdump(err),
			)
		}
	}
}

func TestDefaultAgg(t *testing.T) {
	runTestCases(defaultAggTestCases, t)
}

func TestQuickVar(t *testing.T) {
	runTestCases(quickVarTestCases, t)
}
