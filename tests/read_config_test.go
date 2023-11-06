package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka"
	"github.com/godverv/matreshka/resources"
)

func Test_ReadConfig(t *testing.T) {
	cfgGot, err := matreshka.ParseConfig([]byte(resourcedConfig))
	require.NoError(t, err)

	cfgExpect := &matreshka.AppConfig{
		AppInfo: matreshka.AppInfo{
			Name:            "matreshka",
			Version:         "0.0.1",
			StartupDuration: 10 * time.Second,
		},
		DataSources: []resources.Resource{
			&resources.Postgres{
				Name:    "postgres",
				Host:    "localhost",
				Port:    5432,
				User:    "matreshka",
				Pwd:     "matreshka",
				DbName:  "matreshka",
				SSLMode: "false",
			},
		},
		Server: nil,
	}

	require.Equal(t, cfgExpect, cfgGot)

}
