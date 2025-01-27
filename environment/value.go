package environment

import (
	"reflect"

	"gopkg.in/yaml.v3"
)

type Value struct {
	val typedValue
}

type typedValue interface {
	YamlValue() any
}

func (v Value) MarshalYAML() (interface{}, error) {
	return v.val.YamlValue(), nil
}

func (v *Value) UnmarshalYAML(node *yaml.Node) error {
	return nil
}

func GetType(val any) variableType {
	refV := reflect.ValueOf(val)
	if refV.Kind() == reflect.Ptr {
		refV = refV.Elem()
	}

	refKind := refV.Kind()

	if refKind == reflect.Slice {
		refKind = reflect.TypeOf(val).Elem().Kind()
	}

	return mapReflectTypeToVariableType[refKind]
}
