package tests

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka"
	"github.com/godverv/matreshka/resources"
)

var (
	//go:embed empty_config.yaml
	emptyConfig string
	//go:embed resourced_config.yaml
	resourcedConfig string
)

func Test_CreateEmptyConfig(t *testing.T) {
	cfg := matreshka.NewEmptyConfig()

	bytes, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, emptyConfig, string(bytes))
}

func Test_CreateConfigWithResources(t *testing.T) {
	cfg := matreshka.NewEmptyConfig()
	cfg.DataSources = append(cfg.DataSources, &resources.Postgres{
		AppResource: resources.AppResource{
			ResourceName: "postgres",
		},
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
