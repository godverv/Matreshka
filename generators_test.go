package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/internal/env_parser"
	"github.com/godverv/matreshka/resources"
)

func Test_GenerateGoConfigKeys(t *testing.T) {
	t.Parallel()

	c, err := ParseConfig(fullConfig)
	require.NoError(t, err)

	expected := []env_parser.EnvVal{
		{
			Name:  "matreshka_bool",
			Value: true,
		},
		{
			Name:  "matreshka_duration",
			Value: "10s",
		},
		{
			Name:  "matreshka_int",
			Value: 1,
		},
		{
			Name:  "matreshka_string",
			Value: "not so basic 🤡 string",
		},
		{
			Name: "Resource_postgres",
			Value: &resources.Postgres{
				Name:    "postgres",
				Host:    "localhost",
				Port:    5432,
				User:    "matreshka",
				Pwd:     "matreshka",
				DbName:  "matreshka",
				SSLMode: "false",
			},
		},
		{
			Name: "Api_rest_server",
			Value: &api.Rest{
				Name: "rest_server",
				Port: 8080,
			},
		},
	}

	res, err := GenerateKeys(*c)
	require.NoError(t, err)
	require.Equal(t, expected, res)
}
