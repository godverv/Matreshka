package environment

import (
	"fmt"
	"reflect"
	"strings"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"
	"gopkg.in/yaml.v3"
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
	Value Value        `yaml:"value"`
}

type opt func(*Variable)

func MustNewVariable(name string, val any, opts ...opt) *Variable {
	v, err := NewVariable(name, val, opts...)
	if err != nil {
		panic(err)
	}

	return v
}

func NewVariable(name string, val any, opts ...opt) (*Variable, error) {
	out := &Variable{
		Name: name,
	}

	for _, o := range opts {
		o(out)
	}

	out.Type = toolbox.Coalesce(out.Type, GetType(val))
	if out.Type == "" {
		return nil, errors.Wrap(ErrUnknownEnvVariableType)
	}

	var err error
	out.Value.val, err = mapVariableTypeToTypedValueConstructor[out.Type](val)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return out, nil
}

var (
	mapReflectTypeToVariableType = map[reflect.Kind]variableType{
		reflect.String:  VariableTypeStr,
		reflect.Bool:    VariableTypeBool,
		reflect.Float64: VariableTypeFloat,
		reflect.Float32: VariableTypeFloat,
		reflect.Int:     VariableTypeInt,
		reflect.Int8:    VariableTypeInt,
		reflect.Int16:   VariableTypeInt,
		reflect.Int32:   VariableTypeInt,
		reflect.Int64:   VariableTypeInt,
		reflect.Uint:    VariableTypeInt,
		reflect.Uint8:   VariableTypeInt,
		reflect.Uint16:  VariableTypeInt,
		reflect.Uint32:  VariableTypeInt,
		reflect.Uint64:  VariableTypeInt,
	}
	mapVariableTypeToTypedValueConstructor = map[variableType]func(in any) (typedValue, error){
		VariableTypeStr: toStringValue,
		VariableTypeInt: toIntVariable,
	}
	mapVariableTypeToYamlNodeParser = map[variableType]func(node *yaml.Node) (typedValue, error){
		VariableTypeInt: fromIntNode,
	}
)

func (v *Variable) UnmarshalYAML(node *yaml.Node) error {
	var value, enum *yaml.Node
	_ = enum
	for cIdx := 0; cIdx < len(node.Content); cIdx += 2 {
		fieldName := node.Content[cIdx].Value
		switch fieldName {
		case "name":
			v.Name = node.Content[cIdx+1].Value
		case "type":
			v.Type = variableType(node.Content[cIdx+1].Value)
		case "value", "enum":
			value = node.Content[cIdx+1]
		}
	}

	val, err := mapVariableTypeToYamlNodeParser[v.Type](value)
	v.Value.val = val

	return err
}

//func (v *Variable) MarshalYAML() (any, error) {
//	out := map[string]any{
//		"name": v.Name,
//		"type": v.Type,
//	}
//
//	if len(v.Enum) != 0 {
//		out["enum"] = v.Enum
//	}
//
//	var val any
//
//	switch v.Type {
//	case VariableTypeInt:
//		val = marshalInt(v.Value)
//	default:
//		val = v.Value
//	}
//
//	out["value"] = val
//
//	return out, nil
//}
//
//func (v *Variable) UnmarshalYAML(unmarshal func(a any) error) error {
//	var vals map[string]any
//	err := unmarshal(&vals)
//	if err != nil {
//		return errors.Wrap(err, "error unmarshalling environment variable")
//	}
//
//	v.Name = vals["name"].(string)
//	v.Type = variableType(vals["type"].(string))
//
//	val := vals["value"]
//	if val == nil {
//		return ErrNoValue
//	}
//
//	v.Value, err = extractValue(val, v.Type)
//	if err != nil {
//		return errors.Wrap(err, "error reading value")
//	}
//
//	enum := vals["enum"]
//	if enum != nil {
//		var ok bool
//		v.Enum, ok = enum.([]any)
//		if !ok {
//			return errors.New(fmt.Sprintf("enum expected to be slice, but got %v ", enum))
//		}
//
//		if !isValueInEnum(v.Value, v.Enum) {
//			return errors.New(fmt.Sprintf("value out of enum: `%v` expected to be in %v", v.Value, enum))
//		}
//	}
//
//	return nil
//}

func (v *Variable) UnmarshalEnv(node *evon.Node) error {
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
		tp = &evon.Node{
			Value: VariableTypeStr,
		}
	}

	v.Type = variableType(fmt.Sprint(tp.Value))
	if enum != nil {
		enumVal, err := extractValue(enum.Value, v.Type)
		if err != nil {
			return errors.Wrap(err, "error extracting enum value")
		}

		enumRef := reflect.ValueOf(enumVal)
		if enumRef.Kind() != reflect.Slice {
			return errors.New("expected enum to be slice, but got " + enumRef.Kind().String())
		}

		for i := 0; i < enumRef.Len(); i++ {
			v.Enum = append(v.Enum, enumRef.Index(i).Interface())
		}
	}

	// TODO: remove onto valid
	val, err := extractValue(node.Value, v.Type)
	if err != nil {
		return errors.Wrap(err, "error extracting value")
	}

	v.Value.val, err = mapVariableTypeToTypedValueConstructor[v.Type](val)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (v *Variable) EnumString() string {
	if len(v.Enum) == 0 {
		return ""
	}

	return toStringArray(reflect.ValueOf(v.Enum))
}

func (v *Variable) ValueString() string {
	ref := reflect.ValueOf(v.Value)
	if ref.Kind() == reflect.Slice {
		return toStringArray(ref)
	}

	return fmt.Sprint(v.Value)
}

// deprecated
func extractValue(val any, vType variableType) (out any, err error) {
	switch vType {
	case VariableTypeInt:
		return extractIntVariable(val)
	case VariableTypeStr:
		return extractStringValue(val)
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

func MapVariableToGoType(variable Variable) (typeName string, importName string) {
	switch variable.Type {
	case VariableTypeInt:
		typeName = "int"
	case VariableTypeStr:
		typeName = "string"
	case VariableTypeBool:
		typeName = "bool"
	case VariableTypeFloat:
		typeName = "float64"
	case VariableTypeDuration:
		typeName = "time.Duration"
		importName = "time"
	default:
		return "any", ""
	}

	varRef := reflect.ValueOf(variable.Value)
	if varRef.Kind() == reflect.Slice {
		typeName = "[]" + typeName
	}

	return typeName, importName
}
