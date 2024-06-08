package environment

import (
	"fmt"

	errors "github.com/Red-Sock/trace-errors"
)

var (
	ErrUnknownEnvVariableType = errors.New("unknown environment variable type")
	ErrNoValue                = errors.New("no value for variable")
)

type variableType string

const (
	VariableTypeInt      variableType = "int"
	VariableTypeStr      variableType = "string"
	VariableTypeBool     variableType = "bool"
	VariableTypeFloat    variableType = "float"
	VariableTypeDuration variableType = "duration"
)

type Variable struct {
	Name  string       `yaml:"name"`
	Type  variableType `yaml:"type"`
	Enum  []any        `yaml:"enum,omitempty"`
	Value any          `yaml:"value,omitempty"`
}

func (a *Variable) UnmarshalYAML(unmarshal func(a any) error) error {
	var vals map[string]any
	err := unmarshal(&vals)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling environment variable")
	}

	a.Name = vals["name"].(string)
	a.Type = variableType(vals["type"].(string))

	val := vals["value"]
	if val == nil {
		return ErrNoValue
	}

	switch a.Type {
	case VariableTypeInt:
		a.Value, err = toIntVariable(val)
	case VariableTypeStr:
		a.Value, err = toStringVariable(val)
	case VariableTypeBool:
		a.Value, err = toBool(val)
	case VariableTypeFloat:
		a.Value, err = toFloatVariable(val)
	case VariableTypeDuration:
		a.Value, err = toDuration(val)
	default:
		err = ErrUnknownEnvVariableType
	}

	if err != nil {
		return errors.Wrap(err, "error reading value")
	}

	enum := vals["enum"]
	if enum != nil {
		var ok bool
		a.Enum, ok = enum.([]any)
		if !ok {
			return errors.New(fmt.Sprintf("enum expected to be slice, but got %v ", enum))
		}

		if !isValueInEnum(a.Value, a.Enum) {
			return errors.New(fmt.Sprintf("value out of enum: `%v` expected to be in %v", a.Value, enum))
		}
	}

	return nil
}
