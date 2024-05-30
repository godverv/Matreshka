package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/servers"
)

func Test_Read_AppConfig(t *testing.T) {
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

func Test_Read_DataSourceConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(resourcedConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.Name = "matreshka"

	cfgExpect.DataSources = append(cfgExpect.DataSources,
		getPostgresClientTest(),
		getRedisClientTest(),
		getGRPCClientTest(),
		getTelegramClientTest(),
	)

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_Read_ServerConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(apiConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.Name = "matreshka"
	cfgExpect.Servers = []servers.Api{
		getRestServerTest(),
		getGRPCServerTest(),
	}

	require.Equal(t, cfgExpect, cfgGot)
}
