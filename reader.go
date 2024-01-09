package matreshka

import (
	"context"
	stderrors "errors"
	"os"
	"strings"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka-be/pkg/api/matreshka_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v3"
)

const (
	VervName      = "VERV_NAME"
	VervConfigUrl = "VERV_CONFIG_URL"
)

func NewEmptyConfig() AppConfig {
	return AppConfig{
		AppInfo:     AppInfo{},
		Resources:   make(Resources, 0),
		Servers:     make(Servers, 0),
		Environment: make(map[string]interface{}),
	}
}

func ReadConfigs(pths ...string) (*AppConfig, error) {
	if len(pths) == 0 {
		return nil, nil
	}

	masterConfig, err := readConfig(pths[0])
	if err != nil {
		return nil, errors.Wrap(err, "error reading master config")
	}

	var errs []error
	for _, pth := range pths[1:] {
		slaveConfig, err := readConfig(pth)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "error reading config at %s", pth))
			continue
		}

		masterConfig = MergeConfigs(masterConfig, slaveConfig)
	}

	client := getClient()
	if client != nil {
		masterConfig = MergeConfigs(getViaApi(client, os.Getenv(VervName)), masterConfig)
	}
	masterConfig = MergeConfigs(getViaEnvironment(), masterConfig)

	if len(errs) != 0 {
		return &masterConfig, stderrors.Join(errs...)
	}

	return &masterConfig, nil
}

func ParseConfig(in []byte) (AppConfig, error) {
	a := NewEmptyConfig()

	err := yaml.Unmarshal(in, &a)
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

	for i := range slave.Resources {
		if master.Resources.get(slave.Resources[i].GetName()) == nil {
			master.Resources = append(master.Resources, slave.Resources[i])
		}
	}

	return master
}

func getClient() (client matreshka_api.MatreshkaBeAPIClient) {
	vervConfigUrl := os.Getenv(VervConfigUrl)
	if vervConfigUrl == "" {
		return nil
	}

	dial, err := grpc.Dial(
		vervConfigUrl,
		// TODO VERV-34
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}

	return matreshka_api.NewMatreshkaBeAPIClient(dial)
}

func getViaApi(client matreshka_api.MatreshkaBeAPIClient, projectName string) AppConfig {
	ac := NewEmptyConfig()
	ctx := context.Background()
	cfg, _ := client.GetConfigRaw(ctx, &matreshka_api.GetConfigRaw_Request{ServiceName: projectName})
	if cfg == nil {
		return ac
	}

	conf, _ := ParseConfig([]byte(cfg.Config))

	return conf
}

func getViaEnvironment() AppConfig {
	envConfig := NewEmptyConfig()

	projectName := os.Getenv(VervName)
	if projectName == "" {
		return envConfig
	}

	envConfig.Environment = readEnvironment(projectName)
	return envConfig
}

func readConfig(pth string) (AppConfig, error) {
	f, err := os.Open(pth)
	if err != nil {
		return NewEmptyConfig(), err
	}

	defer f.Close()

	c := NewEmptyConfig()
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return c, errors.Wrap(err, "error decoding config to struct")
	}

	c.Environment = flatten(c.Environment)

	return c, nil
}

func readEnvironment(prefix string) map[string]interface{} {
	environ := os.Environ()
	out := map[string]interface{}{}

	prefix = strings.ToUpper(prefix)
	for _, variable := range environ {
		idx := strings.Index(variable, "=")
		if idx == -1 {
			continue
		}

		name := strings.ToUpper(variable[:idx])

		if !strings.HasPrefix(name, prefix) {
			continue
		}

		out[strings.ToLower(name[len(prefix)+1:])] = variable[idx+1:]
	}
	return out
}

func flatten(in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	for k, v := range in {
		switch t := v.(type) {
		case map[string]interface{}:
			for flatK, flatV := range flatten(t) {
				out[k+"_"+flatK] = flatV
			}
		default:
			out[k] = v
		}
	}

	return out
}
