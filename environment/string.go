package environment

import (
	"fmt"

	errors "github.com/Red-Sock/trace-errors"
)

func toStringValue(in any) (any, error) {
	switch v := in.(type) {
	case string:
		return v, nil
	case []interface{}:
		out := make([]string, 0, len(v))
		for _, val := range v {
			out = append(out, fmt.Sprint(val))
		}

		return out, nil
	default:
		return "", errors.New(fmt.Sprintf("can't convert %T to a string", in))
	}
}
