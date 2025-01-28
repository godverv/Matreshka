package matreshka

import (
	"strings"

	"go.redsock.ru/evon"
	"go.redsock.ru/rerrors"
	"gopkg.in/yaml.v3"
)

var (
	ErrNotFound       = rerrors.New("no such key in config")
	ErrUnexpectedType = rerrors.New("error casting value to target type")
)

type AppConfig struct {
	AppInfo          `yaml:"app_info,omitempty"`
	DataSources      `yaml:"data_sources,omitempty"`
	Servers          `yaml:"servers,omitempty"`
	Environment      `yaml:"environment,omitempty"`
	ServiceDiscovery `yaml:"service_discovery,omitempty"`
}

func (a *AppConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(*a)
}

func (a *AppConfig) Unmarshal(b []byte) error {
	a.DataSources = DataSources{}
	a.Servers = Servers{}
	a.Environment = Environment{}

	err := yaml.Unmarshal(b, a)
	if err != nil {
		return rerrors.Wrap(err)
	}

	envNameReplacer := strings.NewReplacer(" ", evon.ObjectSplitter, evon.FieldSplitter, evon.ObjectSplitter)

	for i := range a.Environment {
		a.Environment[i].Name = envNameReplacer.Replace(a.Environment[i].Name)
	}

	return nil
}
