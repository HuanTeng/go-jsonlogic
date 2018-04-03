package jsonlogic

import (
	"fmt"
)

const (
	defaultAggregator = "_default_aggregator"
	quickVarPrefix    = '$'
	quickVarOp        = "_quick_access"
)

func (jl *jsonLogic) apply(rule RuleType, data DataType) (result DataType, err error) {
	switch rule := rule.(type) {
	case nil, bool, float64, string, []interface{}:
		return DataType(rule), nil
	case map[string]interface{}:
		if len(rule) == 1 {
			for opName, params := range rule {
				op, ok := jl.ops[opName]
				if ok {
					result, err = jl.applyOperator(op, params, data)
				} else {
					if len(opName) > 1 && opName[0] == quickVarPrefix {
						varOp, ok := jl.ops[quickVarOp]
						if !ok {
							return nil, fmt.Errorf("quick access op not defined")
						}
						result, err = jl.applyOperator(varOp, []interface{}{opName[1:], params}, data)
					} else {
						return nil, fmt.Errorf("operator %s not found", opName)
					}
				}
			}
		} else {
			aggOp, ok := jl.ops[defaultAggregator]
			if !ok {
				return nil, fmt.Errorf("multiple keys found but default aggregator not defined")
			}
			result, err = jl.applyOperatorWithParamMap(aggOp, rule, data)
		}
	}

	return result, err
}

func (jl *jsonLogic) applyOperator(op Operator, params interface{}, data DataType) (DataType, error) {
	var paramRules []RuleType

	switch params := params.(type) {
	case []interface{}:
		paramRules = make([]RuleType, len(params))
		for i, param := range params {
			paramRules[i] = RuleType(param)
		}
	default:
		paramRules = make([]RuleType, 1)
		paramRules[0] = RuleType(params)
	}

	return op.Operate(jl, data, paramRules)
}

func (jl *jsonLogic) applyOperatorWithParamMap(op Operator, params map[string]interface{}, data DataType) (DataType, error) {
	var paramRules []RuleType

	paramRules = make([]RuleType, 0, len(params))
	for opName, param := range params {
		paramRules = append(paramRules, RuleType(map[string]interface{}{
			opName: param,
		}))
	}

	return op.Operate(jl, data, paramRules)
}
