package matreshka

import (
	stderrors "errors"
	"os"
	"sort"
	"strings"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/environment"
)

const (
	VervName = "VERV_NAME"
)

func NewEmptyConfig() AppConfig {
	return AppConfig{
		AppInfo:     AppInfo{},
		DataSources: make(DataSources, 0),
		Servers:     make(Servers, 0),
		Environment: make(Environment, 0),
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

	prefix, env := getEnvVars()

	masterEnvStorage := evon.NodeStorage{}
	masterEnv, err := evon.MarshalEnvWithPrefix(prefix, &masterConfig)
	if err != nil {
		return masterConfig, errors.Wrap(err, "error marshalling to env")
	}
	masterEnvStorage.AddNode(masterEnv)

	for _, n := range env {
		masterNode, ok := masterEnvStorage[n.Name]
		if !ok {
			masterEnvStorage[n.Name] = n
		} else {
			masterNode.Value = n.Value
		}
	}
	masterConfig = NewEmptyConfig()
	err = evon.UnmarshalWithNodesAndPrefix(prefix, masterEnvStorage, &masterConfig)
	if err != nil {
		return masterConfig, errors.Wrap(err, "error unmarshalling back to config")
	}

	sort.Slice(masterConfig.Environment, func(i, j int) bool {
		return masterConfig.Environment[i].Name < masterConfig.Environment[j].Name
	})

	return masterConfig, nil
}

func ParseConfig(in []byte) (AppConfig, error) {
	a := NewEmptyConfig()

	err := a.Unmarshal(in)
	if err != nil {
		return a, err
	}

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

	for _, slaveVal := range slave.Environment {
		var mv *environment.Variable
		for _, masterVal := range master.Environment {
			if masterVal.Name == slaveVal.Name {
				mv = masterVal
				break
			}
		}
		if mv != nil {
			continue
		}

		master.Environment = append(master.Environment, slaveVal)
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

	for _, slaveOverride := range slave.ServiceDiscovery.Overrides {
		found := false
		for _, masterOverride := range master.ServiceDiscovery.Overrides {
			if masterOverride.ServiceName == slaveOverride.ServiceName {
				found = true
				break
			}
		}

		if !found {
			master.ServiceDiscovery.Overrides = append(master.ServiceDiscovery.Overrides, slaveOverride)
		}
	}
	return master
}

func getEnvVars() (prefix string, envConfig evon.NodeStorage) {
	envConfig = evon.NodeStorage{}

	projectName := os.Getenv(VervName)
	if projectName == "" {
		return "", envConfig
	}

	environ := os.Environ()

	prefix = strings.ToUpper(projectName)

	for _, variable := range environ {
		idx := strings.Index(variable, "=")
		if idx == -1 {
			continue
		}

		name := strings.ToUpper(variable[:idx])

		if strings.HasPrefix(name, prefix) {
			envConfig.AddNode(&evon.Node{
				Name:       name,
				Value:      os.Getenv(name),
				InnerNodes: nil,
			})
		}
	}

	return prefix, envConfig
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

	return c, nil
}
