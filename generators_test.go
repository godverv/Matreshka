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
			Name: "Resource_redis",
			Value: &resources.Redis{
				Name: "redis",
				Host: "localhost",
				Port: 6379,
				User: "",
				Pwd:  "",
				Db:   0,
			},
		},
		{
			Name: "Resource_telegram",
			Value: &resources.Telegram{
				Name:   "telegram",
				ApiKey: "some_secure_key",
			},
		},
		{
			Name: "Resource_grpc_rscli_example",
			Value: &resources.GRPC{
				Name:             "grpc_rscli_example",
				ConnectionString: "0.0.0.0:50051",
				Module:           "github.com/Red-Sock/rscli_example",
			},
		},
		{
			Name: "Api_rest",
			Value: &api.Rest{
				Name: "rest",
				Port: 8080,
			},
		},
		{
			Name: "Api_grpc",
			Value: &api.GRPC{
				Name: "grpc",
				Port: 50051,
			},
		},
	}

	res, err := GenerateKeys(*c)
	require.NoError(t, err)
	require.Equal(t, expected, res)
}
