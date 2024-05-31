package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/servers"
)

func Test_Read_Config(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(fullConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.AppInfo = AppInfo{
		Name:            "matreshka",
		Version:         "v0.0.1",
		StartupDuration: 10 * time.Second,
	}

	cfgExpect.DataSources = append(cfgExpect.DataSources,
		getPostgresClientTest(),
		getRedisClientTest(),
		getTelegramClientTest(),
		getGRPCClientTest(),
	)

	cfgExpect.Servers = []servers.Api{
		getRestServerTest(),
		getGRPCServerTest(),
	}

	cfgExpect.Environment = map[string]interface{}{
		"matreshka_int":      1,
		"matreshka_string":   "not so basic 🤡 string",
		"matreshka_bool":     true,
		"matreshka_duration": "10s",
	}

	require.Equal(t, cfgExpect, cfgGot)
}
