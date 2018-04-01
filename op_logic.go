package jsonlogic

type opIf struct{}
type opStrictEqual struct{}
type opStrictNEqual struct{}
type opNot struct{}
type opNotNot struct{}
type opOr struct{}
type opAnd struct{}
type opAndBool struct{}

// IsTrue converts supported interface{} into boolean
func IsTrue(v interface{}) bool {
	switch v := v.(type) {
	case nil:
		return false
	case bool:
		return v
	case float64:
		return v != 0
	case string:
		return v != ""
	case []interface{}:
		return len(v) > 0
	default:
		return true
	}
}

func (opIf) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	for {
		if len(params) == 0 {
			return nil, nil
		}
		if len(params) == 1 {
			return applier.Apply(params[0], data)
		}

		// Here len(params) must >= 2
		cond, err := applier.Apply(params[0], data)
		if err != nil {
			return nil, err
		}

		if IsTrue(cond) {
			return applier.Apply(params[1], data)
		} else if len(params) < 2 {
			return nil, nil
		}

		// Here we should do:
		// return op.Operate(applier, data, params[2:])
		// However golang does not support Tail-recursion optimization, so we do it manually
		params = params[2:]
	}
}

func (opStrictEqual) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	var v1, v2 interface{}
	var err error

	if len(params) >= 1 {
		v1, err = applier.Apply(params[0], data)
		if err != nil {
			return nil, err
		}
	}
	if len(params) >= 2 {
		v2, err = applier.Apply(params[1], data)
		if err != nil {
			return nil, err
		}
	}

	return v1 == v2, nil
}

func (opStrictNEqual) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	var v1, v2 interface{}
	var err error

	if len(params) >= 1 {
		v1, err = applier.Apply(params[0], data)
		if err != nil {
			return nil, err
		}
	}
	if len(params) >= 2 {
		v2, err = applier.Apply(params[1], data)
		if err != nil {
			return nil, err
		}
	}

	return v1 != v2, nil
}

func (opNot) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	var v1 interface{}
	var err error
	if len(params) >= 1 {
		v1, err = applier.Apply(params[0], data)
		if err != nil {
			return nil, err
		}
	}

	return !IsTrue(v1), nil
}

func (opNotNot) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	var v1 interface{}
	var err error
	if len(params) >= 1 {
		v1, err = applier.Apply(params[0], data)
		if err != nil {
			return nil, err
		}
	}

	return IsTrue(v1), nil
}

func (opOr) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	var v interface{}
	var err error
	for _, p := range params {
		v, err = applier.Apply(p, data)
		if err != nil {
			return nil, err
		}
		if IsTrue(v) {
			return v, nil
		}
	}
	return v, nil
}

func (opAnd) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	var v interface{}
	var err error
	for _, p := range params {
		v, err = applier.Apply(p, data)
		if err != nil {
			return nil, err
		}
		if !IsTrue(v) {
			return v, nil
		}
	}
	return v, nil
}

func (opAndBool) Operate(applier LogicApplier, data DataType, params []RuleType) (DataType, error) {
	var v interface{}
	var err error
	for _, p := range params {
		v, err = applier.Apply(p, data)
		if err != nil {
			return nil, err
		}
		if !IsTrue(v) {
			return false, nil
		}
	}
	return true, nil
}
