package jsonlogic_test

import (
	"testing"
)

var logicTestCases = []TestCase{
	{
		name:   "true",
		rule:   `{"if" : [ true, "yes", "no" ]}`,
		data:   `{}`,
		expect: "yes",
	},
	{
		name:   "false",
		rule:   `{"if" : [ false, "yes", "no" ]}`,
		data:   `{}`,
		expect: "no",
	},
	{
		name:   "false",
		rule:   `{"if" : [ false, "yes" ]}`,
		data:   `{}`,
		expect: nil,
	},
	{
		name: "elseif",
		rule: `{"if" : [
			{"<": [{"var":"temp"}, 0] }, "freezing",
			{"<": [{"var":"temp"}, 100] }, "liquid",
			"gas"
		  ]}`,
		data:   `{"temp":55}`,
		expect: "liquid",
	},
	{
		name:   "1 === 1",
		rule:   `{"===": [1, 1]}`,
		data:   `{}`,
		expect: true,
	},
	{
		name:   "1 === 2",
		rule:   `{"===": [1, 2]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   `"a" !== "a"`,
		rule:   `{"!==": ["a", "a"]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   `1 !== "1"`,
		rule:   `{"!==": [1, "1"]}`,
		data:   `{}`,
		expect: true,
	},
	{
		name:   `![true]`,
		rule:   `{"!": [true]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   `! true`,
		rule:   `{"!": true}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   `!![[]]`,
		rule:   `{"!!": [[]]}`,
		data:   `{}`,
		expect: false,
	},
	{
		name:   `!!["0"]`,
		rule:   `{"!!": ["0"]}`,
		data:   `{}`,
		expect: true,
	},
	{
		name:   `or`,
		rule:   `{"or":[false, "1", ""]}`,
		data:   `{}`,
		expect: "1",
	},
	{
		name:   `or`,
		rule:   `{"or":[false, 0, ""]}`,
		data:   `{}`,
		expect: "",
	},
	{
		name:   `and`,
		rule:   `{"and":[true,"a",3]}`,
		data:   `{}`,
		expect: float64(3),
	},
	{
		name:   `and`,
		rule:   `{"and":[true,"",3]}`,
		data:   `{}`,
		expect: "",
	},
}

func TestLogicalOp(t *testing.T) {
	runTestCases(logicTestCases, t)
}
