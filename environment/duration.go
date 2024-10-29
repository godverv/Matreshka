package environment

import (
	"fmt"
	"time"

	errors "github.com/Red-Sock/trace-errors"
)

func toDuration(val any) (any, error) {
	switch v := val.(type) {
	case string:
		return time.ParseDuration(v)
	case time.Duration:
		return v, nil
	case []interface{}:
		out := make([]time.Duration, 0, len(v))
		for _, val := range v {
			b, err := toDuration(val)
			if err != nil {
				return nil, errors.Wrap(err, "error extracting duration from array value")
			}

			out = append(out, b.(time.Duration))
		}

		return out, nil

	default:
		return 0, errors.New(fmt.Sprintf("can't cast %T to time.Duration", val))
	}
}
