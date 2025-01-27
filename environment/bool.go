package environment

import (
	"fmt"
	"strconv"

	errors "go.redsock.ru/rerrors"
	"gopkg.in/yaml.v3"
)

type boolValue struct {
	v bool
}

func (s *boolValue) YamlValue() any {
	return s.v
}

type boolSliceValue struct {
	v []bool
}

func (s *boolSliceValue) YamlValue() any {
	return s.v
}

func toBoolValue(val any) (typedValue, error) {
	switch v := val.(type) {
	case string:
		b, err := strconv.ParseBool(v)
		return &boolValue{v: b}, err
	case bool:
		return &boolValue{v: v}, nil
	case []interface{}:
		out := make([]bool, 0, len(v))
		for _, val := range v {
			b, ok := val.(bool)
			if !ok {
				return nil, errors.New(fmt.Sprintf("invalid type for bool array %T", val))
			}

			out = append(out, b)
		}

		return &boolSliceValue{v: out}, nil
	case []bool:
		return &boolSliceValue{v: v}, nil
	default:
		return nil, errors.New(fmt.Sprintf("can't cast %T to bool", val))
	}
}

func fromBoolNode(node *yaml.Node) (typedValue, error) {
	if node.Kind == yaml.ScalarNode {
		b, err := strconv.ParseBool(node.Value)
		return &boolValue{v: b}, err
	}

	if node.Kind == yaml.SequenceNode {
		floatSlice := &boolSliceValue{}

		for _, child := range node.Content {
			b, err := strconv.ParseBool(child.Value)
			if err != nil {
				return nil, errors.Wrap(err, "error parsing bool value from yaml node")
			}

			floatSlice.v = append(floatSlice.v, b)
		}

		return floatSlice, nil
	}

	return nil, errors.New("Expected Bool OR Bool Slice type, got yaml %s", node.Tag)
}

func extractBool(val any) (any, error) {
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
