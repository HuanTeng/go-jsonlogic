package jsonlogic_test

import (
	"testing"
)

var varTestCases = []TestCase{
	{
		name: "string",
		rule: `{"var": ["a", "abc"]}`,
		data: `{"a": {"a1":[12, "ax"], "a2":34}}`,
		expect: map[string]interface{}{
			"a1": []interface{}{
				float64(12),
				"ax",
			},
			"a2": float64(34),
		},
	},
	{
		name:   "number",
		rule:   `{"var": [2, "abc"]}`,
		data:   `[{"a": {"a1":[12, "ax"], "a2":34}, "b": 2}, 2, 3]`,
		expect: float64(3),
	},
	{
		name:   "combined",
		rule:   `{"var": ["0.a.a1.1", "abc"]}`,
		data:   `[{"a": {"a1":[12, "ax"], "a2":34}, "b": 2}, 2, 3]`,
		expect: "ax",
	},
	{
		name:   "default",
		rule:   `{"var": [3, "abc"]}`,
		data:   `[{"a": {"a1":[12, "ax"], "a2":34}, "b": 2}, 2, 3]`,
		expect: "abc",
	},
	{
		name:   "combined default",
		rule:   `{"var": ["0.a.a1.2", "abc"]}`,
		data:   `[{"a": {"a1":[12, "ax"], "a2":34}, "b": 2}, 2, 3]`,
		expect: "abc",
	},
}

func TestVar(t *testing.T) {
	runTestCases(varTestCases, t)
}
