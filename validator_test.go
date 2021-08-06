package jsonlogic_test

import (
	"encoding/json"
	"fmt"
	"github.com/HuanTeng/go-jsonlogic"
	"reflect"
	"testing"
)

var cases = []TestCase{
	{
		name: "valid Rule:0",
		rule: `{
        ">": [
          {
            "var": "variable1.value"
          }]
		}`,
		expect: nil,
	},
	{
		name: "invalid Rule",
		rule: `{
        ">": [
          {
            "variable": "variable1.value"
          }]
		}`,
		expect: fmt.Errorf("invalid operator: variable"),
	},
	{
		name: "Valid Rule:1",
		rule: `{
        "or": {
          ">": [
            {
              "var": "variable1.value"
            },
            {
              "var": "variable1.value"
            }
          ],
          "<": [
            [
              {
                "var": "variable1.value"
              },
              {
                "var": "variable1.value"
              }
            ],
            70
          ]
        }
      }`,
		expect: nil,
	},
	{
		name: "Invalid Rule - 1",
		rule: `{
        "or": {
          "equals": [
            {
              "var": "variable1.value"
            },
            {
              "var": "variable2.value"
            }
          ],
          "fuzzy_match": [
            [
              {
                "var": "variable1.value"
              },
              {
                "var": "variable2.value"
              }
            ],
            70
          ]
        }
      }`,
		expect: fmt.Errorf("invalid operator: equals"),
	},
}

func TestJsonLogic_Validate(t *testing.T) {
	var rule interface{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(c.rule), &rule); err != nil {
				t.Errorf("rule error: %s", err)
			}
			err := jsonlogic.Validate(rule)
			if !reflect.DeepEqual(err, c.expect) {
				t.Errorf("Case %s: expect error %+v got %+v", c.name, c.expect, err)
			}
		})
	}
}
