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

type EnvNode struct {
	Name       string
	Value      any
	InnerNodes []*EnvNode
}

type Unmarshaler interface {
	UnmarshalEnv(env *EnvNode) error
}

type ValueMapFunc func(v *EnvNode) error

func UnmarshalEnv(bytes []byte, in any) error {
	return unmarshal("", bytes, in)
}

func UnmarshalEnvWithPrefix(prefix string, bytes []byte, in any) error {
	return unmarshal(prefix, bytes, in)
}

func ParseDotEnv(bytes []byte) map[string]*EnvNode {

	const root = ""

	nodesMap := map[string]*EnvNode{
		root: {},
	}

	name := ""
	var value any

	start := 0
	for idx := range bytes {
		switch bytes[idx] {
		case '=':
			name = root + "_" + string(bytes[start:idx])
			start = idx + 1
		case '\n':
			value = string(bytes[start:idx])
			start = idx + 1

			node := &EnvNode{
				Name:  name,
				Value: value,
			}
			nodesMap[name] = node

			nameParts := strings.Split(name, "_")
			// todo подумать над пустыми именами
			parentNodePath := nameParts[0]

			for _, namePart := range nameParts[1:] {
				parentNode := nodesMap[parentNodePath]
				if parentNode == nil {
					parentNode = &EnvNode{
						Name: parentNodePath,
					}
					nodesMap[parentNodePath] = parentNode
				}
				currentNodePath := parentNodePath + "_" + namePart

				newNode := &EnvNode{
					Name: currentNodePath,
				}
				if _, ok := nodesMap[newNode.Name]; !ok {
					parentNode.InnerNodes = append(parentNode.InnerNodes, newNode)
					nodesMap[newNode.Name] = newNode
				}
				parentNodePath = currentNodePath
			}
		}
	}

	return nodesMap
}

func unmarshal(prefix string, bytes []byte, in any) error {
	envVals := ParseDotEnv(bytes)

	targetMap := structToMap(prefix, in)

	for key, srcVal := range envVals {
		targetVal, ok := targetMap[key]
		if !ok {
			continue
		}
		_ = targetVal(srcVal)
	}

	return nil
}

func envValToMap(envVals []EnvNode) (map[string]EnvNode, error) {
	fileEnvs := make(map[string]EnvNode)

	for _, e := range envVals {
		if _, ok := fileEnvs[e.Name]; ok {
			return nil, errors.Wrap(ErrDuplicatedEnvKey, e.Name)
		}
		fileEnvs[e.Name] = e
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
		if !v.CanAddr() {
			return
		}
		k := v.Addr()

		//if k.IsNil() {
		//	k.Set(reflect.New(k.Type().Elem()))
		//	k = k.Elem()
		//}
		val := k.Interface()
		customMarshaller, ok := val.(Unmarshaler)
		if !ok {
			panic("Slices of non basic type require sliceMarshaller to be implemented")
		}
		m[strings.ToUpper(prefix)] = customMarshaller.UnmarshalEnv
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
	return func(src *EnvNode) error {
		target.SetString(extractString(reflect.ValueOf(src.Value)))
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
	return func(src *EnvNode) error {
		target.SetInt(extractInt(reflect.ValueOf(src.Value)))
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
	return func(src *EnvNode) error {
		target.SetUint(extractUint(reflect.ValueOf(src.Value)))
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
	return func(src *EnvNode) error {
		target.SetInt(extractDuration(reflect.ValueOf(src.Value)))
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
	return func(src *EnvNode) error {
		target.SetBool(extractBool(reflect.ValueOf(src.Value)))
		return nil
	}
}
