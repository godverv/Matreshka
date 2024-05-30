package env

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	errors "github.com/Red-Sock/trace-errors"
)

var ErrDuplicatedEnvKey = errors.New("environment key collision")

type Unmarshaler interface {
	UnmarshalEnv(env []EnvVal) error
}

type ValueMapFunc func(v reflect.Value) error

func UnmarshalEnv(bytes []byte, in any) error {
	return unmarshal("", bytes, in)
}

func UnmarshalEnvWithPrefix(prefix string, bytes []byte, in any) error {
	return unmarshal(prefix, bytes, in)
}

func ParseDotEnv(bytes []byte) []EnvVal {
	var envVals []EnvVal

	ev := EnvVal{}

	start := 0
	for idx := range bytes {
		switch bytes[idx] {
		case '=':
			ev.Name = string(bytes[start:idx])
			start = idx + 1
		case '\n':
			ev.Value = string(bytes[start:idx])
			start = idx + 1
			envVals = append(envVals, ev)
		}
	}

	return envVals
}

func unmarshal(prefix string, bytes []byte, in any) error {
	envVals := ParseDotEnv(bytes)
	fileEnvs, err := envValToMap(envVals)
	if err != nil {
		return errors.Wrap(err, "error getting map of dotenv")
	}

	targetMap := structToMap(prefix, in)

	for key, srcVal := range fileEnvs {
		targetVal, ok := targetMap[key]
		if !ok {
			continue
		}
		_ = targetVal(reflect.ValueOf(srcVal))
	}

	return nil
}

func envValToMap(envVals []EnvVal) (map[string]any, error) {
	fileEnvs := make(map[string]any)

	for _, e := range envVals {
		if _, ok := fileEnvs[e.Name]; ok {
			return nil, errors.Wrap(ErrDuplicatedEnvKey, e.Name)
		}
		fileEnvs[e.Name] = e.Value
	}

	return fileEnvs, nil
}

func structToMap(prefix string, in any) map[string]ValueMapFunc {
	v := reflect.ValueOf(in)
	m := map[string]ValueMapFunc{}
	reflectValToMap(prefix, v, m)

	return m
}

func reflectValToMap(prefix string, v reflect.Value, m map[string]ValueMapFunc) {
	kind := v.Kind()

	switch kind {
	case reflect.Pointer, reflect.Struct:
		if prefix != "" {
			prefix += "_"
		}
		if kind == reflect.Pointer {
			v = v.Elem()
		}

		for i := 0; i < v.NumField(); i++ {
			tag := v.Type().Field(i).Tag.Get(envTag)
			if tag == "-" {
				continue
			}

			if tag == "" {
				tag = splitToSnake(v.Type().Field(i).Name)
			}

			reflectValToMap(prefix+tag, v.Field(i), m)
		}

	case reflect.String:
		m[strings.ToUpper(prefix)] = mapString(v)
	case reflect.Bool:
		m[strings.ToUpper(prefix)] = mapBool(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tp := v.Type().Name()
		if tp == "Duration" {
			m[strings.ToUpper(prefix)] = mapDuration(v)
		} else {
			m[strings.ToUpper(prefix)] = mapInt(v)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		m[strings.ToUpper(prefix)] = mapUint(v)
	case reflect.Slice:

	default:
		return
	}
}

func extractString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10)
	default:
		return fmt.Sprint(v.Interface())
	}
}
func mapString(target reflect.Value) ValueMapFunc {
	return func(src reflect.Value) error {
		target.SetString(extractString(src))
		return nil
	}
}

func extractInt(v reflect.Value) int64 {
	switch v.Kind() {
	case reflect.String:
		str := v.String()
		d, _ := strconv.ParseInt(str, 10, 64)
		return d
	case reflect.Int:
		return v.Int()
	case reflect.Uint:
		return int64(v.Uint())
	default:
		return 0
	}
}
func mapInt(target reflect.Value) ValueMapFunc {
	return func(v reflect.Value) error {
		target.SetInt(extractInt(v))
		return nil
	}
}

func extractUint(v reflect.Value) uint64 {
	switch v.Kind() {
	case reflect.String:
		str := v.String()
		d, _ := strconv.ParseUint(str, 10, 64)
		return d
	case reflect.Int:
		return uint64(v.Int())
	case reflect.Uint:
		return v.Uint()
	default:
		return 0
	}
}
func mapUint(target reflect.Value) ValueMapFunc {
	return func(src reflect.Value) error {
		target.SetUint(extractUint(src))
		return nil
	}
}

func extractDuration(v reflect.Value) int64 {
	switch v.Kind() {
	case reflect.String:
		str := v.String()
		d, _ := time.ParseDuration(str)
		return int64(d)

	default:
		return 0
	}
}
func mapDuration(target reflect.Value) ValueMapFunc {
	return func(src reflect.Value) error {
		target.SetInt(extractDuration(src))
		return nil
	}
}
func extractBool(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		b, _ := strconv.ParseBool(v.String())
		return b
	default:
		return false
	}
}

func mapBool(target reflect.Value) ValueMapFunc {
	return func(src reflect.Value) error {
		target.SetBool(extractBool(src))
		return nil
	}
}
