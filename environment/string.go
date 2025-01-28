package environment

import (
	"fmt"
	"slices"
	"strings"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"
	"gopkg.in/yaml.v3"
)

type stringValue struct {
	v string
}

func (s *stringValue) Val() any {
	return s.v
}

func (s *stringValue) YamlValue() any {
	return s.v
}

func (s *stringValue) EvonValue() string {
	return s.v
}

type stringSliceValue struct {
	v []string
}

func (s *stringSliceValue) Val() any {
	return s.v
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
}

func (s *stringSliceValue) EvonValue() string {
	return "[" + strings.Join(s.v, ",") + "]"
}

func (s *stringSliceValue) isEnum(val typedValue) error {
	valuesToValidate := make([]string, 0, 1)

	switch actualValue := val.(type) {
	case *stringValue:
		valuesToValidate = append(valuesToValidate, actualValue.v)
	case *stringSliceValue:
		valuesToValidate = append(valuesToValidate, actualValue.v...)
	default:
		return errors.Wrapf(ErrUnexpectedType, "Expected String or String slice but got %T", val)
	}

	for _, valueToValidate := range valuesToValidate {
		if !slices.Contains(s.v, valueToValidate) {
			return errors.Wrapf(ErrEnumValidationFailed, "got %s", valueToValidate)
		}
	}

	return nil
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

func stringValueFromNode(node *yaml.Node) (typedValue, error) {
	if node.Kind == yaml.ScalarNode {
		return &stringValue{v: node.Value}, nil
	}

	if node.Kind == yaml.SequenceNode {
		return stringSliceFromYamlNode(node)
	}

	return nil, errors.New("Unknown yaml type: %s . Expected String or String slice", node.Kind)
}

func stringSliceFromYamlNode(node *yaml.Node) (*stringSliceValue, error) {
	strSlice := &stringSliceValue{}

	for _, child := range node.Content {
		strSlice.v = append(strSlice.v, child.Value)
	}

	return strSlice, nil
}

func stringsSliceFromEvonNode(node *evon.Node) (*stringSliceValue, error) {
	v := node.Value.(string)
	if strings.HasPrefix(v, "[") {
		v = v[1:]
		if strings.HasSuffix(v, "]") {
			v = v[:len(v)-1]
		}
	}

	return &stringSliceValue{v: strings.Split(v, ",")}, nil
}
