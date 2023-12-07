package matreshka

import (
	stderrors "errors"
	"sort"

	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/internal/env_parser"
	"github.com/godverv/matreshka/resources"
)

const (
	apiPrefix      = "Api_"
	resourcePrefix = "Resource_"
)

var ErrNoAppName = stderrors.New("no app name")

func GenerateKeys(c AppConfig) (envs []env_parser.EnvVal, err error) {
	if c.AppInfo.Name == "" {
		return nil, ErrNoAppName
	}

	envs = GenerateEnvironmentKeys(c.AppInfo.Name, c.Environment)
	sort.Slice(envs, func(i, j int) bool {
		return envs[i].Name < envs[j].Name
	})
	envs = append(envs, GenerateResourceConfigKeys(c.Resources...)...)
	envs = append(envs, GenerateApiConfigKeys(c.Servers...)...)

	return envs, nil
}

func GenerateEnvironmentKeys(appName string, in map[string]interface{}) []env_parser.EnvVal {
	return env_parser.ExtractVariables(appName, in)
}

func GenerateResourceConfigKeys(rs ...resources.Resource) []env_parser.EnvVal {
	envs := make([]env_parser.EnvVal, 0, len(rs))
	for idx := range rs {
		envs = append(envs, env_parser.EnvVal{
			Name:  resourcePrefix + rs[idx].GetName(),
			Value: rs[idx],
		})
	}

	return envs
}

func GenerateApiConfigKeys(rs ...api.Api) []env_parser.EnvVal {
	envs := make([]env_parser.EnvVal, 0, len(rs))
	for idx := range rs {
		envs = append(envs, env_parser.EnvVal{
			Name:  apiPrefix + rs[idx].GetName(),
			Value: rs[idx],
		})
	}

	return envs
}
