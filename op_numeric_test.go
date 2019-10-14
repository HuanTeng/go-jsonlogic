package jsonlogic_test

import (
	"testing"
)

var numTestCases = []TestCase{
	{
		name:   "+",
		rule:   `{"+": [6, 2]}`,
		data:   `{}`,
		expect: float64(8),
	},
	{
		name:   "-",
		rule:   `{"-": [6, 2]}`,
		data:   `{}`,
		expect: float64(4),
	},
	{
		name:   "*",
		rule:   `{"*": [6, 2]}`,
		data:   `{}`,
		expect: float64(12),
	},
	{
		name:   "/",
		rule:   `{"/": [6, 2]}`,
		data:   `{}`,
		expect: float64(3),
	},
	{
		name:   "list +",
		rule:   `{"+": [1, 2, 3, 4, 5]}`,
		data:   `{}`,
		expect: float64(15),
	},
	{
		name:   "list *",
		rule:   `{"*": [1, 2, 3, 4, 5]}`,
		data:   `{}`,
		expect: float64(120),
	},
	{
		name:   "inv",
		rule:   `{"-": 2}`,
		data:   `{}`,
		expect: float64(-2),
	},
	{
		name:   "inv neg",
		rule:   `{"-": -2}`,
		data:   `{}`,
		expect: float64(2),
	},
	{
		name:   "to number",
		rule:   `{"+": "3.14"}`,
		data:   `{}`,
		expect: float64(3.14),
	},
	{
		name:   "%",
		rule:   `{"%": [101, 2]}`,
		data:   `{}`,
		expect: float64(1),
	},
	{
		name:   ">",
		rule:   `{">": [2, 1, 0, -1]}`,
		data:   `{}`,
		expect: true,
	},
	{
		name:   "> error input",
		rule:   `{">": ["error", 1, 0, -1]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   "not >",
		rule:   `{">": [2, 2, 1, 0]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   ">=",
		rule:   `{">=": [2, 2, 1, 0]}`,
		data:   `{}`,
		expect: true,
	},
	{
		name:   "not >=",
		rule:   `{">=": [1, 2, 1, 0]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   "<",
		rule:   `{"<": [1, 2, 3, 4]}`,
		data:   `{}`,
		expect: true,
	},
	{
		name:   "not <",
		rule:   `{"<": [1, 2, 2, 4]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   "<=",
		rule:   `{"<=": [1, 2, 2, 4]}`,
		data:   `{}`,
		expect: true,
	},
	{
		name:   "not <=",
		rule:   `{"<=": [1, 2, 1, 4]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   "min",
		rule:   `{"min": [3, 1, 4, 2]}`,
		data:   `{}`,
		expect: float64(1),
	},
	{
		name:   "max",
		rule:   `{"max": [3, 1, 4, 2]}`,
		data:   `{}`,
		expect: float64(4),
	},
}

func TestNumeric(t *testing.T) {
	runTestCases(numTestCases, t)
}
