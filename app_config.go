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
	DataSources `yaml:"data_sources"`
	Servers     `yaml:"server"`
	Environment `yaml:"environment"`
}

func (a *AppConfig) GetAppInfo() AppInfo {
	return a.AppInfo
}

func (a *AppConfig) GetServers() API {
	return &a.Servers
}

func (a *AppConfig) GetDataSources() Resources {
	return &a.DataSources
}

func (a *AppConfig) GetMatreshka() *AppConfig {
	return a
}

func (a *AppConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(*a)
}

func (a *AppConfig) Unmarshal(b []byte) error {
	return yaml.Unmarshal(b, a)
}
