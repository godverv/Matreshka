package environment

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Red-Sock/evon"
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

	a.Value, err = extractValue(val, a.Type)
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

func (a *Variable) UnmarshalEnv(node *evon.Node) error {
	var tp, enum *evon.Node
	for _, n := range node.InnerNodes {
		switch n.Name[len(node.Name)+1:] {
		case "TYPE":
			tp = n
		case "ENUM":
			enum = n
		default:

		}
	}

	if tp == nil {
		return errors.New("environment variable type missing")
	}

	a.Type = variableType(fmt.Sprint(tp.Value))
	if enum != nil {
		enumVal, err := extractValue(enum.Value, a.Type)
		if err != nil {
			return errors.Wrap(err, "error extracting enum value")
		}

		enumRef := reflect.ValueOf(enumVal)
		if enumRef.Kind() != reflect.Slice {
			return errors.New("expected enum to be slice, but got " + enumRef.Kind().String())
		}

		for i := 0; i < enumRef.Len(); i++ {
			a.Enum = append(a.Enum, enumRef.Index(i).Interface())
		}
	}

	var err error
	a.Value, err = extractValue(node.Value, a.Type)
	if err != nil {
		return errors.Wrap(err, "error extracting value")
	}

	return nil
}

func (a *Variable) EnumString() string {
	if len(a.Enum) == 0 {
		return ""
	}

	return toStringArray(reflect.ValueOf(a.Enum))
}

func (a *Variable) ValueString() string {
	ref := reflect.ValueOf(a.Value)
	if ref.Kind() == reflect.Slice {
		return toStringArray(ref)
	}

	return fmt.Sprint(a.Value)
}

func extractValue(val any, vType variableType) (out any, err error) {
	switch vType {
	case VariableTypeInt:
		return toIntVariable(val)
	case VariableTypeStr:
		return toStringValue(val)
	case VariableTypeBool:
		return toBool(val)
	case VariableTypeFloat:
		return toFloatVariable(val)
	case VariableTypeDuration:
		return toDuration(val)
	default:
		return nil, ErrUnknownEnvVariableType
	}
}

func toStringArray(vRef reflect.Value) string {
	vals := make([]string, 0, vRef.Len())
	for i := 0; i < vRef.Len(); i++ {
		vals = append(vals, fmt.Sprint(vRef.Index(i).Interface()))
	}

	return "[" + strings.Join(vals, ",") + "]"
}
