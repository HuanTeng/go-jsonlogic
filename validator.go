package jsonlogic

import (
	"fmt"
)

//TODO: Call operator Specific Validate Methods.
func (jl *jsonLogic) validate(rule RuleType) error {
	switch rule := rule.(type) {
	case nil, bool, float64, string:
		return nil
	case map[string]interface{}:
		for opName, params := range rule {
			_, ok := jl.ops[opName]
			if !ok {
				return fmt.Errorf("invalid operator: " + opName)
			}
			if err := jl.validate(params); err != nil {
				return err
			}
		}
	case []interface{}:
		for _, param := range rule {
			if err := jl.validate(param); err != nil {
				return err
			}
		}

	}

	return nil
}
