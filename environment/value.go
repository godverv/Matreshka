package environment

import (
	"reflect"
)

type Value struct {
	val typedValue
}

func (v Value) Value() any {
	return v.val.Val()
}

type typedValue interface {
	YamlValue() any
	EvonValue() string
	Val() any
}

type typedEnum interface {
	typedValue
	isEnum(value typedValue) error
}

func (v Value) MarshalYAML() (interface{}, error) {
	return v.val.YamlValue(), nil
}

func GetType(val any) variableType {
	refV := reflect.ValueOf(val)

	refKind := refV.Kind()

	if refKind == reflect.Ptr {
		refV = refV.Elem()
		refKind = refV.Kind()
	}

	if refKind == reflect.Slice {
		refKind = reflect.TypeOf(val).Elem().Kind()
	}

	switch refKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		t := refV.Type().String()
		if t == "time.Duration" {
			return VariableTypeDuration
		}
	}

	return mapReflectTypeToVariableType[refKind]
}
