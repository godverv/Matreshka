package matreshka

import (
	"encoding/json"
	"time"

	errors "github.com/Red-Sock/trace-errors"
)

func (a *AppConfig) TryGetInt(key string) (out int, err error) {
	val, ok := a.Environment[key]
	if !ok {
		return 0, errors.Wrap(ErrNotFound, key)
	}

	out, ok = val.(int)
	if !ok {
		return out, errors.Wrapf(ErrUnexpectedType, "wanted: %T actual value %T", out, val)
	}
	return out, nil
}
func (a *AppConfig) GetInt(key string) (out int) {
	out, _ = a.TryGetInt(key)
	return out
}

func (a *AppConfig) TryGetString(key string) (out string, err error) {
	val, ok := a.Environment[key]
	if !ok {
		return "", errors.Wrap(ErrNotFound, key)
	}

	out, ok = val.(string)
	if !ok {
		return out, errors.Wrapf(ErrUnexpectedType, "wanted: %T actual value %T", out, val)
	}
	return out, nil
}
func (a *AppConfig) GetString(key string) (out string) {
	out, _ = a.TryGetString(key)
	return out
}

func (a *AppConfig) TryGetBool(key string) (out bool, err error) {
	val, ok := a.Environment[key]
	if !ok {
		return false, errors.Wrap(ErrNotFound, key)
	}

	out, ok = val.(bool)
	if !ok {
		return out, errors.Wrapf(ErrUnexpectedType, "wanted: %T actual value %T", out, val)
	}
	return out, nil
}
func (a *AppConfig) GetBool(key string) (out bool) {
	out, _ = a.TryGetBool(key)
	return out
}

func (a *AppConfig) TryGetDuration(key string) (out time.Duration, err error) {
	val, ok := a.Environment[key]
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
func (a *AppConfig) GetDuration(key string) (out time.Duration) {
	out, _ = a.TryGetDuration(key)
	return out
}

func (a *AppConfig) TryGetAny(key string) (any, error) {
	val, ok := a.Environment[key]
	if !ok {
		return 0, errors.Wrap(ErrNotFound, key)
	}

	return val, nil
}
func (a *AppConfig) GetAny(key string) any {
	res, _ := a.TryGetAny(key)
	return res
}

func ReadSliceFromConfig[T comparable](cfg *AppConfig, key string, in *[]T) error {
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
