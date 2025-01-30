package matreshka

import (
	stderrors "errors"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strings"

	"go.redsock.ru/evon"
	"go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"

	"go.verv.tech/matreshka/environment"
)

const (
	VervName = "VERV_NAME"
)

func NewEmptyConfig() AppConfig {
	return AppConfig{
		AppInfo:     AppInfo{},
		DataSources: make(DataSources, 0),
		Servers:     make(Servers),
		Environment: make(Environment, 0),
	}
}

func ReadConfigs(paths ...string) (masterConfig AppConfig, err error) {
	masterConfig = NewEmptyConfig()

	if len(paths) != 0 {
		fileConfig, err := getFromFile(paths[0])
		if err != nil {
			return masterConfig, rerrors.Wrap(err, "error reading master config")
		}

		masterConfig = MergeConfigs(masterConfig, fileConfig)

		var errs []error
		for _, pth := range paths[1:] {
			fileConfig, err = getFromFile(pth)
			if err != nil {
				errs = append(errs, rerrors.Wrapf(err, "error reading config at %s", pth))
				continue
			}

			masterConfig = MergeConfigs(masterConfig, fileConfig)
		}

		if len(errs) != 0 {
			return masterConfig, stderrors.Join(errs...)
		}
	}

	prefix, evonStorageFromEnv := getEnvVars(masterConfig.AppInfo)

	//if len(evonStorageFromEnv) != 0 {
	masterConfig, err = mergeWithEnv(prefix, evonStorageFromEnv, masterConfig)
	if err != nil {
		return masterConfig, rerrors.Wrap(err, "error merging config with environment")
	}
	//}

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

	for slavePort, slaveServer := range slave.Servers {
		_, ok := master.Servers[slavePort]
		if !ok {
			master.Servers[slavePort] = slaveServer
		}
	}

	for i := range slave.DataSources {
		if master.DataSources.get(slave.DataSources[i].GetName()) == nil {
			master.DataSources = append(master.DataSources, slave.DataSources[i])
		}
	}

	master.ServiceDiscovery.MakoshUrl =
		toolbox.Coalesce(master.ServiceDiscovery.MakoshUrl, slave.ServiceDiscovery.MakoshUrl)

	master.ServiceDiscovery.MakoshToken =
		toolbox.Coalesce(master.ServiceDiscovery.MakoshToken, slave.ServiceDiscovery.MakoshToken)

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

func getEnvVars(masterInfo AppInfo) (prefixOut string, envConfig evon.NodeStorage) {
	envConfig = evon.NodeStorage{}

	projectName := toolbox.Coalesce(os.Getenv(VervName), masterInfo.ModuleName())

	environ := os.Environ()

	prefixOut = strings.ToUpper(projectName)

	prefixWorking := prefixOut
	// In case project name is specified, appending underscore to clear name from variable
	if prefixWorking != "" {
		prefixWorking += evon.ObjectSplitter
	}

	partsToParse := []string{"APP-INFO", "DATA-SOURCES", "SERVERS", "ENVIRONMENT", "SERVICE-DISCOVERY"}

	for _, variable := range environ {
		idx := strings.Index(variable, "=")
		if idx == -1 {
			continue
		}

		originalName := strings.ToUpper(variable[:idx])

		strippingName := originalName

		if !strings.HasPrefix(strippingName, prefixWorking) {
			continue
		}

		strippingName = strippingName[len(prefixWorking):]
		spaceIndex := strings.Index(strippingName, evon.ObjectSplitter)
		if spaceIndex == -1 {
			continue
		}

		firstPart := strippingName[:spaceIndex]

		if !slices.Contains(partsToParse, firstPart) {
			continue
		}

		envConfig.AddNode(&evon.Node{
			Name:       originalName,
			Value:      os.Getenv(originalName),
			InnerNodes: nil,
		})
	}

	return prefixOut, envConfig
}

func mergeWithEnv(prefix string, evonStorageFromEnv evon.NodeStorage, masterConfig AppConfig) (AppConfig, error) {

	// Storage in Evon format (e.g. object_sub-field-name_leaf-field-name)
	masterEvonStorage := evon.NodeStorage{}
	masterEvonCfg, err := evon.MarshalEnvWithPrefix(prefix, &masterConfig)
	if err != nil {
		return masterConfig, rerrors.Wrap(err, "error marshalling to env")
	}

	// Storage in Basic Environment format (e.g. object_sub_field_name_leaf_field_name)
	masterEnvStorage := map[string]*evon.Node{}
	for key, node := range evonStorageFromEnv {
		key = strings.ReplaceAll(key, evon.FieldSplitter, evon.ObjectSplitter)
		masterEnvStorage[key] = node
	}

	masterEvonStorage.AddNode(masterEvonCfg)

	for _, n := range evonStorageFromEnv {
		masterEvonNode := masterEvonStorage[n.Name]
		if masterEvonNode == nil {
			masterEvonNode = masterEnvStorage[n.Name]
		} else {
			masterEvonStorage.AddNode(n)
		}

		if masterEvonNode != nil {
			masterEvonNode.Value = n.Value
		}
	}

	masterConfig = NewEmptyConfig()
	err = evon.UnmarshalWithNodesAndPrefix(prefix, masterEvonStorage, &masterConfig)
	if err != nil {
		return masterConfig, rerrors.Wrap(err, "error unmarshalling back to config")
	}

	sort.Slice(masterConfig.Environment, func(i, j int) bool {
		return masterConfig.Environment[i].Name < masterConfig.Environment[j].Name
	})

	return masterConfig, nil
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

	fi, err := f.Stat()
	if err != nil {
		return AppConfig{}, rerrors.Wrap(err, "error getting config file info")
	}

	if fi.Size() > 1_000_000 {
		return AppConfig{}, fmt.Errorf("config file too large (more than a 1 MB)")
	}

	c := NewEmptyConfig()

	bts, err := io.ReadAll(f)
	if err != nil {
		return c, rerrors.Wrap(err, "error reading file")
	}

	err = c.Unmarshal(bts)
	if err != nil {
		return c, rerrors.Wrap(err, "error decoding config to struct")
	}

	return c, nil
}
