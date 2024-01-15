package matreshka

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed tests/empty_config.yaml
	emptyConfig []byte
	//go:embed tests/app_config.yaml
	appConfig []byte
	//go:embed tests/resourced_config.yaml
	resourcedConfig []byte
	//go:embed tests/api_config.yaml
	apiConfig []byte
	//go:embed tests/api_half_empty_config.yaml
	apiHalfEmptyConfig []byte
	//go:embed tests/full_config.yaml
	fullConfig []byte
	//go:embed tests/environment_config.yaml
	environmentConfig []byte
	//go:embed tests/invalid_environment_config.yaml
	invalidEnvironmentConfig []byte
)

func Test_CreateEmptyConfig(t *testing.T) {
	t.Parallel()

	cfg := NewEmptyConfig()

	bytes, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, string(emptyConfig), string(bytes))
}

func Test_CreateConfigWithResources(t *testing.T) {
	t.Parallel()

	cfg := NewEmptyConfig()
	cfg.Name = "matreshka"
	cfg.Resources = append(cfg.Resources,
		getPostgresClientTest(),
		getRedisClientTest(),
		getGRPCClientTest(),
		getTelegramClientTest(),
	)

	bytes, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, string(resourcedConfig), string(bytes))
}

func Test_CreateConfigWithServers(t *testing.T) {
	t.Parallel()

	cfg := NewEmptyConfig()
	cfg.Name = "matreshka"
	cfg.Servers = append(cfg.Servers,
		getRestServerTest(),
		getGRPCServerTest(),
	)

	apiMarshalled, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, string(apiConfig), string(apiMarshalled))
}
