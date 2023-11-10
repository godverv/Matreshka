package matreshka

import (
	"time"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	AppInfo     `yaml:"app_info"`
	Resources   `yaml:"data_sources"`
	Servers     `yaml:"server"`
	Environment map[string]interface{} `yaml:"environment"`
}

func (a *AppConfig) TryGetInt(key string) (out int, ok bool) {
	val, ok := a.Environment[key]
	if !ok {
		return 0, false
	}

	out, ok = val.(int)
	return out, ok
}
func (a *AppConfig) GetInt(key string) (out int) {
	out, _ = a.TryGetInt(key)
	return out
}

func (a *AppConfig) TryGetString(key string) (out string, ok bool) {
	val, ok := a.Environment[key]
	if !ok {
		return "", false
	}

	out, ok = val.(string)
	return out, ok
}
func (a *AppConfig) GetString(key string) (out string) {
	out, _ = a.TryGetString(key)
	return out
}

func (a *AppConfig) TryGetBool(key string) (out bool, ok bool) {
	val, ok := a.Environment[key]
	if !ok {
		return false, false
	}

	out, ok = val.(bool)
	return out, ok
}
func (a *AppConfig) GetBool(key string) (out bool) {
	out, _ = a.TryGetBool(key)
	return out
}

func (a *AppConfig) TryGetDuration(key string) (t time.Duration, ok bool) {
	val, ok := a.Environment[key]
	if !ok {
		return 0, false
	}

	timed, ok := val.(string)
	if !ok {
		return 0, false
	}

	t, err := time.ParseDuration(timed)
	if err != nil {
		return 0, false
	}

	return t, ok
}
func (a *AppConfig) GetDuration(key string) (out time.Duration) {
	out, _ = a.TryGetDuration(key)
	return out
}

func (a *AppConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(*a)
}
