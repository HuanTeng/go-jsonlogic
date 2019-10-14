package jsonlogic

import (
	"math"
	"strconv"
)

type opAdd struct{}
type opSub struct{}
type opMul struct{}
type opDiv struct{}
type opMod struct{}
type opGreater struct{}
type opGreaterEqual struct{}
type opLess struct{}
type opLessEqual struct{}
type opMax struct{}
type opMin struct{}

func getFloatNumber(v interface{}) (f float64, ok bool) {
	switch v := v.(type) {
	case nil:
		return 0, ok
	case float64:
		return v, true
	case string:
		if vv, err := strconv.ParseFloat(v, 64); err == nil {
			return vv, true
		}
	}
	return 0, false
}

func binaryOperate(applier LogicApplier, data DataType, params []RuleType, op func(float64, float64) interface{}) (DataType, error) {
	if len(params) <= 1 {
		return nil, nil
	}

	var (
		v1, v2   interface{}
		vv1, vv2 float64
		err      error
		ok       bool
	)

	if v1, err = applier.Apply(params[0], data); err != nil {
		return nil, err
	}
	if vv1, ok = getFloatNumber(v1); !ok {
		return nil, nil
	}

	if v2, err = applier.Apply(params[1], data); err != nil {
		return nil, err
	}
	if vv2, ok = getFloatNumber(v2); !ok {
		return nil, nil
	}

	return op(vv1, vv2), nil
}

func reduceOperate(applier LogicApplier, data DataType, params []RuleType, zero float64, op func(float64, float64) float64) (DataType, error) {
	r := zero

	for _, p := range params {
		v, err := applier.Apply(p, data)
		if err != nil {
			return nil, err
		}
		vv, ok := getFloatNumber(v)
		if !ok {
			return nil, nil
		}
		r = op(r, vv)
		if math.IsNaN(r) {
			break
		}
	}

	return r, nil
}

func (opAdd) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	return reduceOperate(applier, data, params, float64(0), func(v1, v2 float64) float64 {
		return v1 + v2
	})
}

func (opMul) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	return reduceOperate(applier, data, params, float64(1), func(v1, v2 float64) float64 {
		return v1 * v2
	})
}

func (opSub) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	if len(params) == 1 {
		v, err := applier.Apply(params[0], data)
		if err != nil {
			return nil, err
		}

		if vv, ok := getFloatNumber(v); ok {
			return -vv, nil
		}
		return nil, nil
	}

	return binaryOperate(applier, data, params, func(v1, v2 float64) interface{} {
		return v1 - v2
	})
}

func (opDiv) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	return binaryOperate(applier, data, params, func(v1, v2 float64) interface{} {
		return v1 / v2
	})
}

func (opMod) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	return binaryOperate(applier, data, params, func(v1, v2 float64) interface{} {
		return math.Mod(v1, v2)
	})
}

func (opGreater) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	ret, err := reduceOperate(applier, data, params, math.Inf(+1), func(v1, v2 float64) float64 {
		if v1 > v2 {
			return v2
		}
		return math.NaN()
	})
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return false, nil
	}
	return !math.IsNaN(ret.(float64)), nil
}

func (opGreaterEqual) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	ret, err := reduceOperate(applier, data, params, math.Inf(+1), func(v1, v2 float64) float64 {
		if v1 >= v2 {
			return v2
		}
		return math.NaN()
	})
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return false, nil
	}
	return !math.IsNaN(ret.(float64)), nil
}

func (opLess) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	ret, err := reduceOperate(applier, data, params, math.Inf(-1), func(v1, v2 float64) float64 {
		if v1 < v2 {
			return v2
		}
		return math.NaN()
	})
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return false, nil
	}
	return !math.IsNaN(ret.(float64)), nil
}

func (opLessEqual) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	ret, err := reduceOperate(applier, data, params, math.Inf(-1), func(v1, v2 float64) float64 {
		if v1 <= v2 {
			return v2
		}
		return math.NaN()
	})
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return false, nil
	}
	return !math.IsNaN(ret.(float64)), nil
}

func (opMax) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	return reduceOperate(applier, data, params, math.Inf(-1), func(v1, v2 float64) float64 {
		if v1 > v2 {
			return v1
		}
		return v2
	})
}

func (opMin) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	return reduceOperate(applier, data, params, math.Inf(+1), func(v1, v2 float64) float64 {
		if v1 < v2 {
			return v1
		}
		return v2
	})
}
