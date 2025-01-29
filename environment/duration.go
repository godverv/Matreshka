package environment

import (
	"fmt"
	"strings"
	"time"

	errors "go.redsock.ru/rerrors"
	"gopkg.in/yaml.v3"
)

type durationValue struct {
	v time.Duration
}

func (s *durationValue) Val() any {
	return s.v
}

func (s *durationValue) EvonValue() string {
	return s.v.String()
}

func (s *durationValue) YamlValue() any {
	return s.v
}

type durationSliceValue struct {
	v []time.Duration
}

func (s *durationSliceValue) Val() any {
	return s.v
}

func (s *durationSliceValue) EvonValue() string {
	strs := make([]string, 0, len(s.v))

	for _, d := range s.v {
		strs = append(strs, d.String())
	}

	return "[" + strings.Join(strs, ",") + "]"
}

func (s *durationSliceValue) YamlValue() any {
	node := &yaml.Node{
		Kind:  yaml.SequenceNode,
		Style: yaml.FlowStyle,
	}

	for _, b := range s.v {
		node.Content = append(node.Content, &yaml.Node{
			Kind:  yaml.ScalarNode,
			Tag:   "!!str",
			Value: b.String(),
		})
	}

	return node
}

func toDuration(val any) (typedValue, error) {
	switch v := val.(type) {
	case string:
		d, err := time.ParseDuration(v)
		return &durationValue{v: d}, err

	case time.Duration:
		return &durationValue{v: v}, nil

	case []time.Duration:
		return &durationSliceValue{v: v}, nil

	default:
		return nil, errors.New(fmt.Sprintf("can't cast %T to time.Duration", val))
	}
}

func fromDurationNode(val *yaml.Node) (v typedValue, err error) {
	if val.Kind == yaml.ScalarNode {
		dur, err := time.ParseDuration(val.Value)
		return &durationValue{v: dur}, err
	}

	if val.Kind == yaml.SequenceNode {
		arr := make([]time.Duration, len(val.Content))
		for idx := range val.Content {
			arr[idx], err = time.ParseDuration(val.Content[idx].Value)
			if err != nil {
				return nil, errors.Wrap(err)
			}
		}

		return &durationSliceValue{v: arr}, nil
	}

	return nil, nil
}
