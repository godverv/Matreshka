package matreshka

import (
	stderrors "errors"
	"os"

	errors "github.com/Red-Sock/trace-errors"
	"go.deanishe.net/env"
	"gopkg.in/yaml.v3"
)

func NewEmptyConfig() *AppConfig {
	return &AppConfig{
		AppInfo:     AppInfo{},
		Resources:   make(Resources, 0),
		Servers:     make(Servers, 0),
		Environment: make(map[string]interface{}),
	}
}

func ReadConfig(pth string) (*AppConfig, error) {
	f, err := os.Open(pth)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	c := NewEmptyConfig()
	err = yaml.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding config to struct")
	}

	return c, nil
}

func ReadConfigs(pths ...string) (*AppConfig, error) {
	if len(pths) == 0 {
		return nil, nil
	}

	masterConfig, err := ReadConfig(pths[0])
	if err != nil {
		return nil, errors.Wrap(err, "error reading master config")
	}

	var errs []error
	for _, pth := range pths[1:] {
		slaveConfig, err := ReadConfig(pth)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "error reading config at %s", pth))
			continue
		}

		MergeConfigs(masterConfig, slaveConfig)
	}
	envConfig := NewEmptyConfig()
	err = env.Bind(envConfig)
	if err != nil {
		errs = append(errs, errors.Wrap(err, "error obtaining env config"))
	}

	MergeConfigs(envConfig, masterConfig)

	if len(errs) != 0 {
		return masterConfig, stderrors.Join(errs...)
	}

	return masterConfig, nil
}

func ParseConfig(in []byte) (*AppConfig, error) {
	a := NewEmptyConfig()
	err := yaml.Unmarshal(in, a)
	if err != nil {
		return nil, err
	}

	namedMap := make(map[string]interface{})

	for k, v := range a.Environment {
		namedMap[a.Name+"_"+k] = v
	}

	a.Environment = namedMap

	return a, nil
}

func MergeConfigs(master, slave *AppConfig) {
	if master.Name == "" {
		master.Name = slave.Name
	}
	if master.Version == "" {
		master.Version = slave.Version
	}
	if master.StartupDuration == 0 {
		master.StartupDuration = slave.StartupDuration
	}

	for name, value := range slave.Environment {
		if _, ok := master.Environment[name]; !ok {
			master.Environment[name] = value
		}
	}

	for i := range slave.Servers {
		if master.Servers.get(slave.Servers[i].GetName()) == nil {
			master.Servers = append(master.Servers, slave.Servers[i])
		}
	}

	for i := range slave.Resources {
		if master.Resources.get(slave.Resources[i].GetName()) == nil {
			master.Resources = append(master.Resources, slave.Resources[i])
		}
	}
}
