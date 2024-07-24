package labels

import (
	"fmt"
	"strconv"
)

type Requirement struct {
	key      string
	operator Operator
	values   []string
}

type Operator string

const (
	Exists         Operator = "exists"
	NotExist       Operator = "!"
	Equals         Operator = "="
	NotEquals      Operator = "!="
	In             Operator = "in"
	NotIn          Operator = "notin"
	GreaterThan    Operator = "gt"
	LessThan       Operator = "lt"
	GreaterOrEqual Operator = "ge"
	LessOrEqual    Operator = "le"
)

// NewRequirement creates a new Requirement based on the provided key, operator and values.
// If any of these rules is violated, an error is returned when constructed.
// 1. The Operator must be one of the predefined constants.
// such as: Exists, NotExist, Equals, NotEquals, In, NotIn, GreaterThan, LessThan, GreaterOrEqual, LessOrEqual
// 2. For operators GreaterThan, LessThan, GreaterOrEqual, LessOrEqual, the values must be a number.
// 3. For operators In, NotIn, the values must not be empty.
// 4. For operators Exists, NotExist,  the values must not be empty.
// 5. For operators Equals, NotEquals, the values must have exactly one entry.
func NewRequirement(key string, op Operator, vals []string) (*Requirement, error) {
	if key == "" {
		return nil, fmt.Errorf("key must not be empty")
	}
	switch op {
	case Exists, NotExist:
		if len(vals) != 0 {
			return nil, fmt.Errorf("values must be empty for operator %q", op)
		}
	case Equals, NotEquals:
		if len(vals) != 1 {
			return nil, fmt.Errorf("values must have exactly one entry for operator %q", op)
		}
	case In, NotIn:
		if len(vals) == 0 {
			return nil, fmt.Errorf("values must not be empty for operator %q", op)
		}
	case GreaterThan, LessThan, GreaterOrEqual, LessOrEqual:
		if len(vals) != 1 {
			return nil, fmt.Errorf("values must have exactly one entry for operator %q", op)
		}
		if _, err := strconv.ParseInt(vals[0], 10, 64); err != nil {
			return nil, fmt.Errorf("invalid value for operator %q: %v", op, err)
		}
	default:
		return nil, fmt.Errorf("unknown operator %q", op)
	}
	return &Requirement{key: key, operator: op, values: vals}, nil
}

func (r *Requirement) hasValue(value string) bool {
	for _, s := range r.values {
		if s == value {
			return true
		}
	}
	return false
}

func (r *Requirement) Matches(labels Labels) bool {
	switch r.operator {
	case Equals, In:
		if !labels.Has(r.key) {
			return false
		}
		return r.hasValue(labels.Get(r.key))
	case NotEquals, NotIn:
		if !labels.Has(r.key) {
			return true
		}
		return !r.hasValue(labels.Get(r.key))
	case Exists:
		return labels.Has(r.key)
	case NotExist:
		return !labels.Has(r.key)
	case GreaterThan, LessThan, GreaterOrEqual, LessOrEqual:
		if !labels.Has(r.key) {
			return false
		}
		matchValue, err := strconv.ParseInt(labels.Get(r.key), 10, 64)
		if err != nil {
			return false
		}
		requireValue, err := strconv.ParseInt(r.values[0], 10, 64)
		if err != nil {
			return false
		}
		return (r.operator == GreaterThan && matchValue > requireValue) || (r.operator == LessThan && matchValue < requireValue) ||
			(r.operator == GreaterOrEqual && matchValue >= requireValue) || (r.operator == LessOrEqual && matchValue <= requireValue)
	default:
		return false
	}
}
