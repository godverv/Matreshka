package environment

import (
	"fmt"
	"strconv"
	"strings"

	errors "go.redsock.ru/rerrors"
)

func toFloatVariable(val any) (any, error) {
	switch switchValue := val.(type) {
	case []interface{}:
		v, err := anySliceToFloatSlice(switchValue)
		return v, err
	case string:
		if switchValue[0] == '[' && switchValue[len(switchValue)-1] == ']' {
			strSlice := strings.Split(switchValue[1:len(switchValue)-1], ",")
			anySlice := make([]any, 0, len(switchValue)/2)
			for _, v := range strSlice {
				anySlice = append(anySlice, v)
			}
			return anySliceToFloatSlice(anySlice)
		}
		return anyToFloat(val)
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
	case string:
		return strconv.ParseFloat(switchValue, 64)
	default:
		return 0, errors.New(fmt.Sprintf("can't cast %T to float", val))
	}
}
