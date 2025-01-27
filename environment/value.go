package environment

import (
	"reflect"
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
