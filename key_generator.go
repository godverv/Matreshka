package matreshka

import (
	stderrors "errors"
	"sort"

	errors "github.com/Red-Sock/trace-errors"

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

	envs = GenerateEnvironmentKeys(c.Environment)
	sort.Slice(envs, func(i, j int) bool {
		return envs[i].Name < envs[j].Name
	})

	resourcesEnvs, err := GenerateResourceConfigKeys(c.Resources)
	if err != nil {
		return nil, errors.Wrap(err, "error extracting resource env keys")
	}

	envs = append(envs, resourcesEnvs...)

	apiEnvs, err := GenerateApiConfigKeys(c.Servers)
	if err != nil {
		return nil, errors.Wrap(err, "error extracting api env keys")
	}

	envs = append(envs, apiEnvs...)

	return envs, nil
}

func GenerateEnvironmentKeys(in map[string]interface{}) []env_parser.EnvVal {
	return env_parser.ExtractVariables("", in)
}

func GenerateResourceConfigKeys(rs []resources.Resource) ([]env_parser.EnvVal, error) {
	envs := make([]env_parser.EnvVal, 0, len(rs))
	for idx := range rs {
		key := resourcePrefix + rs[idx].GetName()
		envs = append(envs, env_parser.EnvVal{
			Name:  key,
			Value: rs[idx],
		})

		resourceEnvs, err := env_parser.ExtractFromAny(key, rs[idx])
		if err != nil {
			return nil, errors.Wrap(err, "error extracting resource values")
		}

		envs = append(envs, resourceEnvs...)
	}

	return envs, nil
}

func GenerateApiConfigKeys(rs []api.Api) ([]env_parser.EnvVal, error) {
	envs := make([]env_parser.EnvVal, 0, len(rs))
	for idx := range rs {
		key := apiPrefix + rs[idx].GetName()
		envs = append(envs, env_parser.EnvVal{
			Name:  key,
			Value: rs[idx],
		})

		apiEnvs, err := env_parser.ExtractFromAny(key, rs[idx])
		if err != nil {
			return nil, errors.Wrap(err, "error extracting server values")
		}
		envs = append(envs, apiEnvs...)
	}

	return envs, nil
}
