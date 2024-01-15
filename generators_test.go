package matreshka

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/internal/env_parser"
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
			Name:  "Resource_postgres",
			Value: getPostgresClientTest(),
		},
		{
			Name:  "Resource_redis",
			Value: getRedisClientTest(),
		},
		{
			Name:  "Resource_telegram",
			Value: getTelegramClientTest(),
		},
		{
			Name:  "Resource_grpc_rscli_example",
			Value: getGRPCClientTest(),
		},
		{
			Name:  "Api_rest_server",
			Value: getRestServerTest(),
		},
		{
			Name:  "Api_grpc_server",
			Value: getGRPCServerTest(),
		},
	}

	res, err := GenerateKeys(c)
	require.NoError(t, err)
	sort.Slice(res, func(i, j int) bool {
		return res[i].Name < res[i].Name
	})

	require.Equal(t, expected, res)
}
