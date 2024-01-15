package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/api"
)

func Test_ReadAppConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(appConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.AppInfo = AppInfo{
		Name:            "matreshka",
		Version:         "0.0.1",
		StartupDuration: 10 * time.Second,
	}

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_ReadDataSourceConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(resourcedConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.Name = "matreshka"

	cfgExpect.Resources = append(cfgExpect.Resources,
		getPostgresClientTest(),
		getRedisClientTest(),
		getGRPCClientTest(),
		getTelegramClientTest(),
	)

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_ReadApiConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(apiConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.Name = "matreshka"
	cfgExpect.Servers = []api.Api{
		getRestServerTest(),
		getGRPCServerTest(),
	}

	require.Equal(t, cfgExpect, cfgGot)
}
