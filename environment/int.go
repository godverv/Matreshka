package environment

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	errors "go.redsock.ru/rerrors"
)

type intValue struct {
	v int
}

func (v *intValue) YamlValue() any {
	return v.v
}

type intSliceValue struct {
	v []int
}

func (v *intSliceValue) YamlValue() any {
	ranges := make([]string, 0, len(v.v))
	sort.Slice(v.v, func(i, j int) bool {
		return v.v[i] < v.v[j]
	})

	if len(v.v) == 0 {
		return "[]"
	}

	convertToRange := func(start, end int) string {
		newRange := strconv.Itoa(start)
		if start != end {
			newRange += "-" + strconv.Itoa(end)
		}

		return newRange
	}

	prev := v.v[0]
	rangeStart := prev

	for _, v := range v.v[1:] {
		if v-prev != 1 {
			ranges = append(ranges, convertToRange(rangeStart, prev))

			prev = v
			rangeStart = v
		}
		prev = v
	}

	ranges = append(ranges, convertToRange(rangeStart, prev))

	return "[" + strings.Join(ranges, ",") + "]"
}

func toIntVariable(val any) (typedValue, error) {
	switch switchValue := val.(type) {
	case string:
		if switchValue[0] == '[' {
			v, err := extractIntSliceFromString(switchValue)
			return &intSliceValue{v: v}, err
		}

		//TODO test [1-20, 30] MUST fail
		rangeSeparator := strings.Index(switchValue, "-")
		if rangeSeparator != -1 {
			v, err := extractIntRange(rangeSeparator, switchValue)
			return &intSliceValue{v: v}, err
		}

		v, err := strconv.Atoi(switchValue)
		return &intValue{v: v}, err

	case []interface{}:
		v, err := anySliceToIntSlice(switchValue)
		return &intSliceValue{v: v}, err

	default:
		v, err := anyToInt(val)
		return &intValue{v: v}, err
	}
}

func extractIntVariable(val any) (any, error) {
	switch switchValue := val.(type) {
	case string:
		if switchValue[0] == '[' {
			return extractIntSliceFromString(switchValue)
		}

		rangeSeparator := strings.Index(switchValue, "-")
		if rangeSeparator != -1 {
			return extractIntRange(rangeSeparator, switchValue)
		}

		return strconv.Atoi(switchValue)

	case []interface{}:
		return anySliceToIntSlice(switchValue)

	default:
		return anyToInt(val)
	}
}

func extractIntSliceFromString(switchValue string) ([]int, error) {
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
				rng, err := extractIntRange(rangeSeparator, v)
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
	case int8:
		return int(switchValue), nil
	case int16:
		return int(switchValue), nil
	case int32:
		return int(switchValue), nil
	case int64:
		return int(switchValue), nil
	case uint:
		return int(switchValue), nil
	case uint8:
		return int(switchValue), nil
	case uint16:
		return int(switchValue), nil
	case uint32:
		return int(switchValue), nil
	case uint64:
		return int(switchValue), nil
	default:
		return 0, errors.New(fmt.Sprintf("can't cast %T to int", val))
	}
}

func extractIntRange(rangeSeparatorIdx int, strValue string) ([]int, error) {
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

func marshalInt(in any) string {
	switch newIn := in.(type) {
	case []int:
		ranges := make([]string, 0, len(newIn))
		sort.Slice(newIn, func(i, j int) bool {
			return newIn[i] < newIn[j]
		})

		if len(newIn) == 0 {
			return "[]"
		}

		convertToRange := func(start, end int) string {
			newRange := strconv.Itoa(start)
			if start != end {
				newRange += "-" + strconv.Itoa(end)
			}

			return newRange
		}

		prev := newIn[0]
		rangeStart := prev

		for _, v := range newIn[1:] {
			if v-prev != 1 {
				ranges = append(ranges, convertToRange(rangeStart, prev))

				prev = v
				rangeStart = v
			}
			prev = v
		}

		ranges = append(ranges, convertToRange(rangeStart, prev))

		return "[" + strings.Join(ranges, ",") + "]"
	default:
		return fmt.Sprint(newIn)
	}
}
