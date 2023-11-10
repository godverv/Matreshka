package matreshka

import (
	"time"

	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"
)

var (
	ErrNotFound = errors.New("no such key in config")
	ErrParsing  = errors.New("error casting value to wanted type")
)

type AppConfig struct {
	AppInfo     `yaml:"app_info"`
	Resources   `yaml:"data_sources"`
	Servers     `yaml:"server"`
	Environment map[string]interface{} `yaml:"environment"`
}

func (a *AppConfig) TryGetInt(key string) (out int, err error) {
	val, ok := a.Environment[key]
	if !ok {
		return 0, errors.Wrap(ErrNotFound, key)
	}

	out, ok = val.(int)
	if !ok {
		return out, errors.Wrapf(ErrParsing, "wanted: %T actual value %T", out, val)
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
		return out, errors.Wrapf(ErrParsing, "wanted: %T actual value %T", out, val)
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
		return out, errors.Wrapf(ErrParsing, "wanted: %T actual value %T", out, val)
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
		return 0, errors.Wrap(ErrParsing, "error parsing value to string before parsing duration")
	}

	out, err = time.ParseDuration(timed)
	if err != nil {
		return 0, errors.Wrapf(ErrParsing, "error parssing duration")
	}

	return out, nil
}
func (a *AppConfig) GetDuration(key string) (out time.Duration) {
	out, _ = a.TryGetDuration(key)
	return out
}

func (a *AppConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(*a)
}
