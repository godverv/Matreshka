package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/environment"
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

	cfgExpect.Environment = []environment.Variable{
		{
			Name:  "database_max_connections",
			Type:  environment.VariableTypeInt,
			Value: 1,
		},
		{
			Name:  "welcome_string",
			Type:  environment.VariableTypeStr,
			Value: "not so basic 🤡 string",
		},
		{
			Name:  "one of welcome string",
			Type:  environment.VariableTypeStr,
			Value: "one",
			Enum:  []any{"one", "two", "three"},
		},
		{
			Name:  "true falser",
			Type:  environment.VariableTypeBool,
			Value: true,
		},
		{
			Name:  "request timeout",
			Type:  environment.VariableTypeDuration,
			Value: time.Second * 10,
		},
		{
			Name:  "available ports",
			Type:  environment.VariableTypeInt,
			Value: []int{10, 12, 34, 35, 36, 37, 38, 39, 40},
		},
		{
			Name:  "usernames to ban",
			Type:  environment.VariableTypeStr,
			Value: []string{"hacker228", "mothe4acker"},
		},
		{
			Name:  "credit percent",
			Type:  environment.VariableTypeFloat,
			Value: 0.01,
		},
		{
			Name:  "credit percents based on year of birth",
			Type:  environment.VariableTypeFloat,
			Value: []float64{0.01, 0.02, 0.03, 0.04},
		},
	}

	require.Equal(t, cfgExpect, cfgGot)
}
