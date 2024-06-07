package environment

import (
	"fmt"
	"strconv"

	errors "github.com/Red-Sock/trace-errors"
)

func toBool(val any) (bool, error) {
	switch v := val.(type) {
	case string:
		return strconv.ParseBool(v)
	case bool:
		return v, nil
	default:
		return false, errors.New(fmt.Sprintf("can't cast %T to bool", val))
	}
}
