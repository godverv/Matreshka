package environment

import (
	"fmt"
	"strconv"
	"strings"

	errors "github.com/Red-Sock/trace-errors"
)

func toIntVariable(val any) (any, error) {
	switch switchValue := val.(type) {
	case string:
		if switchValue[0] == '[' {
			return stringToIntSlice(switchValue)
		}

		rangeSeparator := strings.Index(switchValue, "-")
		if rangeSeparator != -1 {
			return toIntRange(rangeSeparator, switchValue)
		}

		return strconv.Atoi(switchValue)

	case []interface{}:
		return anySliceToIntSlice(switchValue)

	default:
		return anyToInt(val)
	}
}

func stringToIntSlice(switchValue string) ([]int, error) {
	separatedVals := strings.Split(switchValue[1:len(switchValue)-1], ",")

	anyVals := make([]any, 0, len(separatedVals))
	for _, v := range separatedVals {
		anyVals = append(anyVals, v)
	}

	return anySliceToIntSlice(anyVals)
}

func anySliceToIntSlice(value []any) ([]int, error) {
	out := make([]int, 0, len(value))

	for _, v := range value {
		switch v := v.(type) {
		case string:
			rangeSeparator := strings.Index(v, "-")
			if rangeSeparator != -1 {
				rng, err := toIntRange(rangeSeparator, v)
				if err != nil {
					return nil, errors.Wrap(err, "error converting value to int range")
				}

				out = append(out, rng...)
			} else {
				newInt, err := strconv.Atoi(v)
				if err != nil {
					return nil, errors.Wrap(err, "error converting value to int")
				}

				out = append(out, newInt)
			}
		default:
			val, err := anyToInt(v)
			if err != nil {
				return nil, errors.Wrap(err, "error converting any to int")
			}

			out = append(out, val)
		}
	}

	return out, nil
}

func anyToInt(val any) (int, error) {
	switch switchValue := val.(type) {
	case int:
		return switchValue, nil

	default:
		return 0, errors.New(fmt.Sprintf("can't cast %T to int", val))
	}
}

func toIntRange(rangeSeparatorIdx int, strValue string) ([]int, error) {
	firstNumber, err := strconv.Atoi(strValue[:rangeSeparatorIdx])
	if err != nil {
		return nil, errors.Wrap(err, "error parsing first number of range to int")
	}

	secondNumber, err := strconv.Atoi(strValue[rangeSeparatorIdx+1:])
	if err != nil {
		return nil, errors.Wrap(err, "error parsing second number of range to int")
	}

	out := make([]int, 0, secondNumber-firstNumber)

	for i := firstNumber; i <= secondNumber; i++ {
		out = append(out, i)
	}

	return out, nil
}
