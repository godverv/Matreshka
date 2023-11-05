package matreshka

import (
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/api"
)

type AppConfig struct {
	AppInfo     AppInfo   `yaml:"app_info"`
	DataSources Resources `yaml:"data_sources,omitempty"`
	Server      []api.Api `yaml:"server,omitempty"`
}

func (a *AppConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(*a)
}
