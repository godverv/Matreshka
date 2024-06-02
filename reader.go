package matreshka

import (
	stderrors "errors"
	"os"
	"strings"

	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"
)

const (
	VervName = "VERV_NAME"
	ApiURL   = "MATRESHKA_URL"
)

func NewEmptyConfig() AppConfig {
	return AppConfig{
		AppInfo:     AppInfo{},
		DataSources: make(DataSources, 0),
		Servers:     make(Servers, 0),
		Environment: make(Environment),
	}
}

func ReadConfigs(paths ...string) (AppConfig, error) {
	masterConfig := NewEmptyConfig()

	if len(paths) != 0 {
		fileConfig, err := getFromFile(paths[0])
		if err != nil {
			return masterConfig, errors.Wrap(err, "error reading master config")
		}

		masterConfig = MergeConfigs(masterConfig, fileConfig)

		var errs []error
		for _, pth := range paths[1:] {
			fileConfig, err = getFromFile(pth)
			if err != nil {
				errs = append(errs, errors.Wrapf(err, "error reading config at %s", pth))
				continue
			}

			masterConfig = MergeConfigs(masterConfig, fileConfig)
		}

		if len(errs) != 0 {
			return masterConfig, stderrors.Join(errs...)
		}
	}

	masterConfig = MergeConfigs(getFromEnvironment(), masterConfig)

	return masterConfig, nil
}

func ParseConfig(in []byte) (AppConfig, error) {
	a := NewEmptyConfig()

	err := a.Unmarshal(in)
	if err != nil {
		return a, err
	}

	a.Environment = flatten(a.Environment)

	namedMap := make(map[string]interface{})

	for k, v := range a.Environment {
		namedMap[a.Name+"_"+k] = v
	}

	a.Environment = namedMap

	return a, nil
}

func MergeConfigs(master, slave AppConfig) AppConfig {
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

	for i := range slave.DataSources {
		if master.DataSources.get(slave.DataSources[i].GetName()) == nil {
			master.DataSources = append(master.DataSources, slave.DataSources[i])
		}
	}

	return master
}

func getFromEnvironment() AppConfig {
	envConfig := NewEmptyConfig()

	projectName := os.Getenv(VervName)
	if projectName == "" {
		return envConfig
	}

	environ := os.Environ()

	prefix := strings.ToUpper(projectName)
	for _, variable := range environ {
		idx := strings.Index(variable, "=")
		if idx == -1 {
			continue
		}

		name := strings.ToUpper(variable[:idx])

		if !strings.HasPrefix(name, prefix) {
			continue
		}

		envConfig.Environment[strings.ToLower(name[len(prefix)+1:])] = variable[idx+1:]
	}

	return envConfig
}

func getFromFile(pth string) (AppConfig, error) {
	f, err := os.Open(pth)
	if err != nil {
		return NewEmptyConfig(), err
	}

	defer func() {
		closerErr := f.Close()
		if err == nil {
			err = closerErr
			return
		}

		err = stderrors.Join(err, closerErr)
	}()

	c := NewEmptyConfig()
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return c, errors.Wrap(err, "error decoding config to struct")
	}

	c.Environment = flatten(c.Environment)

	return c, nil
}
