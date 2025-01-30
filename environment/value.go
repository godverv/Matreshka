package environment

import (
	"reflect"
)

type Value struct {
	val typedValue
}

type typedValue interface {
	YamlValue() any
	EvonValue() string
	Val() any
}

func (v Value) Value() any {
	if v.val != nil {
		return v.val.Val()
	}

	return nil
}

func (v Value) MarshalYAML() (interface{}, error) {
	if v.val != nil {
		return v.val.YamlValue(), nil
	}

	return nil, nil
}

func (v Value) String() string {
	return v.val.EvonValue()
}

func GetType(val any) variableType {
	refV := reflect.ValueOf(val)

	refKind := refV.Kind()

	if refKind == reflect.Ptr {
		refV = refV.Elem()
		refKind = refV.Kind()
	}
	isSlice := false
	if refKind == reflect.Slice {
		refKind = reflect.TypeOf(val).Elem().Kind()
		isSlice = true
	}

	switch refKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var nativeType string
		if isSlice {
			nativeType = reflect.TypeOf(val).Elem().String()
		} else {
			nativeType = refV.Type().String()
		}

		if nativeType == "time.Duration" {
			return VariableTypeDuration
		}
	}

	return mapReflectTypeToVariableType[refKind]
}
