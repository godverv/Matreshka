package environment

import (
	"fmt"
	"time"

	errors "github.com/Red-Sock/trace-errors"
)

func toDuration(val any) (time.Duration, error) {
	switch v := val.(type) {
	case string:
		return time.ParseDuration(v)
	default:
		return 0, errors.New(fmt.Sprintf("can't cast %T to time.Duration", val))
	}
}
