package matreshka

import (
	"os"

	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"
)

func NewEmptyConfig() *AppConfig {
	return &AppConfig{}
}

func ReadConfig(pth string) (*AppConfig, error) {
	f, err := os.Open(pth)
	if err != nil {
		return nil, err
	}

	defer f.Close()

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
