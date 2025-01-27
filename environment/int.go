package environment

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	errors "go.redsock.ru/rerrors"
	"gopkg.in/yaml.v3"
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

func mp(inMapped []int) string {
	if len(inMapped) == 0 {
		return "[]"
	}

	ranges := make([]string, 0, len(inMapped))

	if slices.IsSorted(inMapped) {

	}

	return "[" + strings.Join(ranges, ",") + "]"
}

func (v *intSliceValue) YamlValue() any {
	intRanges := make([]string, 0, len(v.v))

	if slices.IsSorted(v.v) {
		return v.asYamlRange()
	}

	node := &yaml.Node{
		Kind:  yaml.SequenceNode,
		Style: yaml.FlowStyle,
	}

	if len(v.v) == 0 {
		return node
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
			intRanges = append(intRanges, convertToRange(rangeStart, prev))

			prev = v
			rangeStart = v
		}
		prev = v
	}

	intRanges = append(intRanges, convertToRange(rangeStart, prev))

	for _, r := range intRanges {
		node.Content = append(node.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: r,
			})
	}

	return node
}

func (v *intSliceValue) asYamlRange() *yaml.Node {
	node := &yaml.Node{
		Kind:  yaml.SequenceNode,
		Style: yaml.FlowStyle,
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
			node.Content = append(node.Content, &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: convertToRange(rangeStart, prev),
			})

			prev = v
			rangeStart = v
		}
		prev = v
	}

	node.Content = append(node.Content, &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: convertToRange(rangeStart, prev),
	})

	return node
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
	case []int:
		return &intSliceValue{v: switchValue}, nil
	case []int8:
		return &intSliceValue{v: toIntSlice(switchValue)}, nil
	case []int16:
		return &intSliceValue{v: toIntSlice(switchValue)}, nil
	case []int32:
		return &intSliceValue{v: toIntSlice(switchValue)}, nil
	case []int64:
		return &intSliceValue{v: toIntSlice(switchValue)}, nil
	default:
		v, err := anyToInt(val)
		return &intValue{v: v}, err
	}
}

func fromIntNode(node *yaml.Node) (typedValue, error) {
	if node.Kind == yaml.ScalarNode {
		i, err := strconv.Atoi(node.Value)
		return &intValue{i}, err
	}

	if node.Kind == yaml.SequenceNode {
		intSlice := &intSliceValue{}

		for _, child := range node.Content {

			i := 0
			var err error

			switch child.Tag {
			case "!!int":
				i, err = strconv.Atoi(child.Value)
				if err != nil {
					return nil, errors.Wrap(err, "could not parse int: %s ", child.Value)
				}

				intSlice.v = append(intSlice.v, i)
			case "!!str":
				minusIndex := strings.Index(child.Value, "-")
				if minusIndex == -1 || minusIndex == 0 && strings.Count(child.Value, "-") < 2 {
					i, err = strconv.Atoi(child.Value)
					if err != nil {
						return nil, errors.Wrap(err, "could not parse string as int: %s ", child.Value)
					}

					intSlice.v = append(intSlice.v, i)
				} else {

					minusIndex = strings.Index(child.Value[1:], "-") + 1
					var firstInt, lastInt int
					firstInt, err = strconv.Atoi(child.Value[:minusIndex])
					if err != nil {
						return nil, errors.Wrap(err, "error parsing first int in sequence")
					}

					lastInt, err = strconv.Atoi(child.Value[minusIndex+1:])
					if err != nil {
						return nil, errors.Wrap(err, "error parsing last int in sequence")
					}

					for ; firstInt <= lastInt; firstInt++ {
						intSlice.v = append(intSlice.v, firstInt)
					}
				}
			}
		}

		return intSlice, nil
	}

	return nil, errors.New("Expected Int OR Int Slice type, got yaml %s", node.Tag)
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

func toIntSlice[T int | int8 | int16 | int32 | int64](v []T) []int {
	out := make([]int, 0, len(v))
	for _, v := range v {
		out = append(out, int(v))
	}

	return out
}
