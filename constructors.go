package matreshka

import (
	"os"

	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/resources"
	"github.com/godverv/matreshka/server"
)

func NewEmptyConfig() *AppConfig {
	return &AppConfig{
		Server:      make([]server.Server, 0),
		DataSources: make([]resources.Resource, 0),
	}
}

func ReadConfig(pth string) (*AppConfig, error) {
	f, err := os.Open(pth)
	if err != nil {
		return nil, err
	}

	c := &AppConfig{}

	err = yaml.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding config to struct")
	}

	return c, nil
}

func ParseConfig(in []byte) (*AppConfig, error) {
	var a AppConfig
	return &a, yaml.Unmarshal(in, &a)
}
