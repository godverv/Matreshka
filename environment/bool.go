package environment

import (
	"fmt"
	"strconv"

	errors "github.com/Red-Sock/trace-errors"
)

func toBool(val any) (any, error) {
	switch v := val.(type) {
	case string:
		return strconv.ParseBool(v)
	case bool:
		return v, nil
	case []interface{}:
		out := make([]bool, 0, len(v))
		for _, val := range v {
			b, ok := val.(bool)
			if !ok {
				return nil, errors.New(fmt.Sprintf("invalid type for bool array %T", val))
			}

			out = append(out, b)
		}

		return out, nil

	default:
		return false, errors.New(fmt.Sprintf("can't cast %T to bool", val))
	}
}
