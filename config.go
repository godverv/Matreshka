package matreshka

import (
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	AppInfo     AppInfo   `yaml:"app_info"`
	DataSources Resources `yaml:"data_sources,omitempty"`
	Server      Servers   `yaml:"server,omitempty"`
}

func (a *AppConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(*a)
}
