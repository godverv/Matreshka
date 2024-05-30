package matreshka

import (
	stderrors "errors"
	"sort"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka/data_sources"
	"github.com/godverv/matreshka/internal/env"
	"github.com/godverv/matreshka/servers"
)

const (
	apiPrefix      = "Api_"
	resourcePrefix = "Resource_"
)

var ErrNoAppName = stderrors.New("no app name")

func GenerateKeys(c AppConfig) (envs []env.EnvVal, err error) {
	if c.AppInfo.Name == "" {
		return nil, ErrNoAppName
	}

	envs = GenerateEnvironmentKeys(c.Environment)
	sort.Slice(envs, func(i, j int) bool {
		return envs[i].Name < envs[j].Name
	})

	resourcesEnvs, err := GenerateResourceConfigKeys(c.DataSources)
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

func GenerateEnvironmentKeys(in map[string]interface{}) []env.EnvVal {
	return env.ExtractVariables("", in)
}

func GenerateResourceConfigKeys(rs []data_sources.Resource) ([]env.EnvVal, error) {
	envs := make([]env.EnvVal, 0, len(rs))
	for idx := range rs {
		key := resourcePrefix + rs[idx].GetName()
		envs = append(envs, env.EnvVal{
			Name:  key,
			Value: rs[idx],
		})

		resourceEnvs, err := env.ExtractFromAny(key, rs[idx])
		if err != nil {
			return nil, errors.Wrap(err, "error extracting resource values")
		}

		envs = append(envs, resourceEnvs...)
	}

	return envs, nil
}

func GenerateApiConfigKeys(rs []servers.Api) ([]env.EnvVal, error) {
	envs := make([]env.EnvVal, 0, len(rs))
	for idx := range rs {
		key := apiPrefix + rs[idx].GetName()
		envs = append(envs, env.EnvVal{
			Name:  key,
			Value: rs[idx],
		})

		apiEnvs, err := env.ExtractFromAny(key, rs[idx])
		if err != nil {
			return nil, errors.Wrap(err, "error extracting server values")
		}
		envs = append(envs, apiEnvs...)
	}

	return envs, nil
}
