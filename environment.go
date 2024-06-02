package matreshka

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
)

type Environment map[string]any

func (a Environment) TryGetInt(key string) (out int, err error) {
	val, ok := a[key]
	if !ok {
		return 0, errors.Wrap(ErrNotFound, key)
	}

	switch val.(type) {
	case int:
		out, ok = val.(int)
	case string:
		s, _ := val.(string)
		out, err = strconv.Atoi(s)
		ok = err == nil
	default:
		ok = false
	}

	if !ok {
		return out, errors.Wrapf(ErrUnexpectedType, "wanted: int but got %v of type %T", val, val)
	}

	return out, nil
}
func (a Environment) GetInt(key string) (out int) {
	out, _ = a.TryGetInt(key)
	return out
}

func (a Environment) TryGetString(key string) (out string, err error) {
	val, ok := a[key]
	if !ok {
		return "", errors.Wrap(ErrNotFound, key)
	}

	out, ok = val.(string)
	if !ok {
		return out, errors.Wrapf(ErrUnexpectedType, "wanted: %T actual value %T", out, val)
	}
	return out, nil
}
func (a Environment) GetString(key string) (out string) {
	out, _ = a.TryGetString(key)
	return out
}

func (a Environment) TryGetBool(key string) (out bool, err error) {
	val, ok := a[key]
	if !ok {
		return false, errors.Wrap(ErrNotFound, key)
	}

	switch v := val.(type) {
	case bool:
		out = v

	case string:
		v = strings.ToLower(v)
		out = v == "true"
		ok = v == "true" || v == "false"
	default:
		ok = false
	}

	if ok {
		return out, nil
	}

	return out, errors.Wrapf(ErrUnexpectedType, "wanted: string or bool, got: %T", val)
}
func (a Environment) GetBool(key string) (out bool) {
	out, _ = a.TryGetBool(key)
	return out
}

func (a Environment) TryGetDuration(key string) (out time.Duration, err error) {
	val, ok := a[key]
	if !ok {
		return 0, errors.Wrap(ErrNotFound, key)
	}

	timed, ok := val.(string)
	if !ok {
		return 0, errors.Wrap(ErrUnexpectedType, "error parsing value to string before parsing duration")
	}

	out, err = time.ParseDuration(timed)
	if err != nil {
		return 0, errors.Wrapf(ErrUnexpectedType, "error parssing duration")
	}

	return out, nil
}
func (a Environment) GetDuration(key string) (out time.Duration) {
	out, _ = a.TryGetDuration(key)
	return out
}

func (a Environment) TryGetAny(key string) (any, error) {
	val, ok := a[key]
	if !ok {
		return 0, errors.Wrap(ErrNotFound, key)
	}

	return val, nil
}
func (a Environment) GetAny(key string) any {
	res, _ := a.TryGetAny(key)
	return res
}

func (a Environment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(a)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling environment variables")
	}

	newMap := flatten(a)

	for k := range a {
		delete(a, k)
	}

	for k, v := range newMap {
		a[k] = v
	}
	return nil
}

func (a Environment) MarshalEnv(prefix string) []evon.Node {
	if prefix != "" {
		prefix += "_"
	}

	out := make([]evon.Node, 0, len(a))
	for k, v := range a {
		out = append(out, evon.Node{
			Name:  prefix + strings.ToUpper(k),
			Value: v,
		})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Name > out[j].Name
	})

	return out
}
func (a Environment) UnmarshalEnv(rootNode *evon.Node) error {
	for _, n := range rootNode.InnerNodes {
		a[strings.ToLower(n.Name[len(rootNode.Name)+1:])] = n.Value
	}

	return nil
}

func ReadSliceFromConfig[T comparable](cfg AppConfig, key string, in *[]T) error {
	res, ok := cfg.Environment[key]
	if !ok {
		return errors.Wrap(ErrNotFound, key)
	}

	bts, err := json.Marshal(res)
	if err != nil {
		return errors.Wrap(err, "error marshalling value")
	}

	err = json.Unmarshal(bts, in)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling value")
	}

	return nil
}

func flatten(in map[string]any) map[string]any {
	out := make(map[string]any)

	for k, v := range in {
		switch t := v.(type) {
		case Environment:
			for flatK, flatV := range flatten(t) {
				out[k+"_"+flatK] = flatV
			}
		default:
			out[k] = v
		}
	}

	return out
}
