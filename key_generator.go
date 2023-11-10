package matreshka

import (
	stderrors "errors"
	"sort"

	"github.com/godverv/matreshka/internal/env_parser"
)

var ErrNoAppName = stderrors.New("no app name")

func GenerateEnvironmentKeys(c AppConfig) (envs []env_parser.EnvVal, err error) {
	if c.AppInfo.Name == "" {
		return nil, ErrNoAppName
	}

	envs = env_parser.ExtractVariables(c.AppInfo.Name, c.Environment)

	sort.Slice(envs, func(i, j int) bool {
		return envs[i].Name < envs[j].Name
	})

	return envs, nil
}
