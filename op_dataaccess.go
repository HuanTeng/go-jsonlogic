package jsonlogic

import (
	"strconv"
	"strings"
)

type opVar struct{}
type opVarStrictEqual struct{}

func (opVar) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	// No params, return whole data
	if len(params) == 0 {
		return data, nil
	}

	key, err := applier.Apply(params[0], data)
	if err != nil {
		return nil, err
	}

	var keyList []string
	notFound := false

	switch key := key.(type) {
	case string:
		keyList = strings.Split(key, ".")
	case float64:
		keyList = []string{strconv.FormatFloat(key, 'f', -1, 64)}
	default:
		keyList = nil
		notFound = true
	}

	var d interface{} = data
	for _, key := range keyList {
		if notFound {
			break
		}
		switch dv := d.(type) {
		case map[string]interface{}:
			if v, ok := dv[key]; ok {
				d = v
			} else {
				notFound = true
			}
		case []interface{}:
			if kv, err := strconv.ParseUint(key, 10, 64); err == nil && int(kv) < len(dv) {
				d = dv[kv]
			} else {
				notFound = true
			}
		}
	}
	if notFound {
		if len(params) >= 2 {
			d, err = applier.Apply(params[1], data)
			if err != nil {
				return nil, err
			}
		} else {
			d = nil
		}
	}

	return d, nil
}

func (opVarStrictEqual) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	if len(params) < 2 {
		return nil, nil
	}
	result, err := (opVar{}).Operate(applier, data, []RuleType{params[0]})
	if err != nil {
		return nil, err
	}
	result, err = (opStrictEqual{}).Operate(applier, data, []RuleType{result, params[1]})
	return result, err
}
