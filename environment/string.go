package environment

import (
	"fmt"
	"strings"

	errors "go.redsock.ru/rerrors"
	"gopkg.in/yaml.v3"
)

type stringValue struct {
	v string
}

func (s *stringValue) YamlValue() any {
	return s.v
}

type stringSliceValue struct {
	v []string
}

func (s *stringSliceValue) YamlValue() any {
	node := &yaml.Node{
		Kind:  yaml.SequenceNode,
		Style: yaml.FlowStyle,
	}

	node.Content = make([]*yaml.Node, 0, len(s.v))

	for _, v := range s.v {
		node.Content = append(node.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!str",
				Value: v,
			},
		)
	}

	return node

	return "[" + strings.Join(s.v, ",") + "]"
}

func toStringValue(in any) (typedValue, error) {
	switch v := in.(type) {
	case string:
		if v[0] == '[' && v[len(v)-1] == ']' {
			out := strings.Split(v[1:len(v)-1], ",")
			return &stringSliceValue{
				v: out,
			}, nil
		}

		return &stringValue{v: v}, nil
	case []interface{}:
		out := make([]string, 0, len(v))
		for _, val := range v {
			out = append(out, fmt.Sprint(val))
		}

		return &stringSliceValue{v: out}, nil
	case []string:
		return &stringSliceValue{v: v}, nil
	default:
		return nil, errors.New(fmt.Sprintf("can't convert %T to a string", in))
	}
}

func extractStringValue(in any) (any, error) {
	switch v := in.(type) {
	case string:
		if v[0] == '[' && v[len(v)-1] == ']' {
			out := strings.Split(v[1:len(v)-1], ",")
			return out, nil
		}

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
