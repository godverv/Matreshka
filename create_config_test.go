package matreshka

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/resources"
)

var (
	//go:embed tests/empty_config.yaml
	emptyConfig string
	//go:embed tests/app_config.yaml
	appConfig string
	//go:embed tests/resourced_config.yaml
	resourcedConfig string
	//go:embed tests/api_config.yaml
	apiConfig string
	//go:embed tests/full_config.yaml
	fullConfig string
	//go:embed tests/environment_config.yaml
	environmentConfig string
	//go:embed tests/invalid_environment_config.yaml
	invalidEnvironmentConfig string
)

func Test_CreateEmptyConfig(t *testing.T) {
	cfg := NewEmptyConfig()

	bytes, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, emptyConfig, string(bytes))
}

func Test_CreateConfigWithResources(t *testing.T) {
	cfg := NewEmptyConfig()

	cfg.Resources = append(cfg.Resources, &resources.Postgres{
		Name:    "postgres",
		Host:    "localhost",
		Port:    5432,
		DbName:  "matreshka",
		User:    "matreshka",
		Pwd:     "matreshka",
		SSLMode: "false",
	})

	bytes, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, resourcedConfig, string(bytes))
}

func Test_CreateConfigWithServers(t *testing.T) {
	cfg := NewEmptyConfig()
	cfg.Servers = append(cfg.Servers, &api.Rest{
		Name: "rest_server",
		Port: 8080,
	})

	bytes, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, apiConfig, string(bytes))
}
