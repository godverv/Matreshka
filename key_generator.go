package matreshka

import (
	stderrors "errors"
	"sort"

	"github.com/godverv/matreshka/internal/env_parser"
)

var ErrNoAppName = stderrors.New("no app name")

func GenerateKeys(c AppConfig) (envs []env_parser.EnvVal, err error) {
	if c.AppInfo.Name == "" {
		return nil, ErrNoAppName
	}

	envs = env_parser.ExtractVariables(c.AppInfo.Name, c.Environment)

	sort.Slice(envs, func(i, j int) bool {
		return envs[i].Name < envs[j].Name
	})

	for idx := range c.Resources {
		envs = append(envs, env_parser.EnvVal{
			Name:  "Resource_" + c.Resources[idx].GetName(),
			Value: c.Resources[idx],
		})
	}

	for idx := range c.Servers {
		envs = append(envs, env_parser.EnvVal{
			Name:  "Api_" + c.Servers[idx].GetName(),
			Value: c.Servers[idx],
		})
	}

	return envs, nil
}
