package environment

import (
	"fmt"
	"strconv"

	errors "github.com/Red-Sock/trace-errors"
)

func toFloatVariable(val any) (any, error) {
	switch switchValue := val.(type) {
	case []interface{}:
		v, err := anySliceToFloatSlice(switchValue)
		return v, err

	default:
		return anyToFloat(val)
	}
}

func anySliceToFloatSlice(value []any) ([]float64, error) {
	out := make([]float64, 0, len(value))

	for _, v := range value {
		switch v := v.(type) {
		case string:

			newFloat, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, errors.Wrap(err, "error converting value to float")
			}

			out = append(out, newFloat)

		default:
			val, err := anyToFloat(v)
			if err != nil {
				return nil, errors.Wrap(err, "error converting any to float")
			}

			out = append(out, val)
		}
	}

	return out, nil
}

func anyToFloat(val any) (float64, error) {
	switch switchValue := val.(type) {
	case float64:
		return switchValue, nil

	default:
		return 0, errors.New(fmt.Sprintf("can't cast %T to int", val))
	}
}
