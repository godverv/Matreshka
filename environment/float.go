package environment

import (
	"fmt"
	"strconv"
	"strings"

	errors "go.redsock.ru/rerrors"
	"gopkg.in/yaml.v3"
)

type floatValue struct {
	v float64
}

func (v *floatValue) Val() any {
	return v.v
}

func (v *floatValue) EvonValue() string {
	return floatToString(v.v)
}

func (v *floatValue) YamlValue() any {
	return v.v
}

type floatSliceValue struct {
	v []float64
}

func (v *floatSliceValue) Val() any {
	return v.v
}

func (v *floatSliceValue) EvonValue() string {
	outStr := make([]string, 0, len(v.v))

	for _, e := range v.v {
		outStr = append(outStr, floatToString(e))
	}
	return "[" + strings.Join(outStr, ",") + "]"
}

func (v *floatSliceValue) YamlValue() any {

	node := &yaml.Node{
		Kind:  yaml.SequenceNode,
		Style: yaml.FlowStyle,
	}

	if len(v.v) == 0 {
		return node
	}

	for _, r := range v.v {
		node.Content = append(node.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: floatToString(r),
			})
	}

	return node
}

func toFloatVariable(val any) (typedValue, error) {
	toFloat64 := func(in float32) float64 {
		return float64(int(float64(in)*100)) / 100
	}

	switch switchValue := val.(type) {
	case float64:
		return &floatValue{v: switchValue}, nil
	case float32:
		return &floatValue{v: toFloat64(switchValue)}, nil
	case []interface{}:
		v, err := anySliceToFloatSlice(switchValue)
		return &floatSliceValue{v: v}, err
	case string:
		if switchValue[0] == '[' && switchValue[len(switchValue)-1] == ']' {
			strSlice := strings.Split(switchValue[1:len(switchValue)-1], ",")
			anySlice := make([]any, 0, len(switchValue)/2)
			for _, v := range strSlice {
				anySlice = append(anySlice, v)
			}
			v, err := anySliceToFloatSlice(anySlice)
			return &floatSliceValue{v: v}, err
		}

		v, err := anyToFloat(val)
		return &floatValue{v}, err
	case []float64:
		return &floatSliceValue{v: switchValue}, nil
	case []float32:
		newSlice := make([]float64, 0, len(switchValue))
		for _, v := range switchValue {
			newSlice = append(newSlice, toFloat64(v))
		}

		return &floatSliceValue{v: newSlice}, nil
	default:
		v, err := anyToFloat(val)
		return &floatValue{v}, err
	}
}

func fromFloatNode(node *yaml.Node) (typedValue, error) {
	if node.Kind == yaml.ScalarNode {
		i, err := strconv.ParseFloat(node.Value, 64)
		return &floatValue{v: i}, err
	}

	if node.Kind == yaml.SequenceNode {
		floatSlice := &floatSliceValue{}

		for _, child := range node.Content {
			v, err := strconv.ParseFloat(child.Value, 64)
			if err != nil {
				return nil, errors.Wrap(err, "error parsing float value from yaml node")
			}
			floatSlice.v = append(floatSlice.v, v)
		}

		return floatSlice, nil
	}

	return nil, errors.New("Expected Float OR Float Slice type, got yaml %s", node.Tag)
}

func extractFloatVariable(val any) (any, error) {
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

func floatToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
