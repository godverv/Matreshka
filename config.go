package matreshka

import (
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/server"
)

type AppConfig struct {
	AppInfo     AppInfo         `yaml:"app_info"`
	DataSources Resources       `yaml:"data_sources,omitempty"`
	Server      []server.Server `yaml:"server,omitempty"`
}

func (a *AppConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(*a)
}
