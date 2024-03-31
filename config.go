package matreshka

import (
	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"
)

var (
	ErrNotFound       = errors.New("no such key in config")
	ErrUnexpectedType = errors.New("error casting value to wanted type")
)

type AppConfig struct {
	AppInfo     `yaml:"app_info"`
	Resources   `yaml:"data_sources"`
	Servers     `yaml:"server"`
	Environment map[string]interface{} `yaml:"environment"`
}

func (a *AppConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(*a)
}

func (a *AppConfig) Unmarshal(b []byte) error {
	return yaml.Unmarshal(b, a)
}
