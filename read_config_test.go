package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/resources"
)

func Test_ReadAppConfig(t *testing.T) {
	cfgGot, err := ParseConfig([]byte(appConfig))
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
	cfgGot, err := ParseConfig([]byte(resourcedConfig))
	require.NoError(t, err)

	cfgExpect := &AppConfig{}
	cfgExpect.Resources = append(cfgExpect.Resources, &resources.Postgres{
		Name:    "postgres",
		Host:    "localhost",
		Port:    5432,
		User:    "matreshka",
		Pwd:     "matreshka",
		DbName:  "matreshka",
		SSLMode: "false",
	})

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_ReadApiConfig(t *testing.T) {
	cfgGot, err := ParseConfig([]byte(apiConfig))
	require.NoError(t, err)

	cfgExpect := &AppConfig{
		Servers: []api.Api{
			&api.Rest{
				Name: "rest_server",
				Port: 8080,
			},
		},
	}

	require.Equal(t, cfgExpect, cfgGot)
}
