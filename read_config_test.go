package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/resources"
)

func Test_ReadAppConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(appConfig)
	require.NoError(t, err)

	cfgExpect := &AppConfig{
		AppInfo: AppInfo{
			Name:            "matreshka",
			Version:         "0.0.1",
			StartupDuration: 10 * time.Second,
		}}

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_ReadDataSourceConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig([]byte(resourcedConfig))
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.Resources = append(cfgExpect.Resources,
		&resources.Postgres{
			Name:    "postgres",
			Host:    "localhost",
			Port:    5432,
			User:    "matreshka",
			Pwd:     "matreshka",
			DbName:  "matreshka",
			SSLMode: "false",
		},
		&resources.Redis{
			Name: "redis",
			Host: "localhost",
			Port: 6379,
			User: "",
			Pwd:  "",
			Db:   0,
		},
		&resources.GRPC{
			Name:             "grpc_rscli_example",
			ConnectionString: "0.0.0.0:50051",
			Module:           "github.com/Red-Sock/rscli_example",
		},
		&resources.Telegram{
			Name:   "telegram",
			ApiKey: "some_api_key",
		},
	)

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_ReadApiConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(apiConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.Servers = []api.Api{
		&api.Rest{
			Name: "rest_server",
			Port: 8080,
		},
		&api.GRPC{
			Name: "grpc_server",
			Port: 50051,
		},
	}

	require.Equal(t, cfgExpect, cfgGot)
}
