package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/internal/env"
	"github.com/godverv/matreshka/resources"
)

func Test_marshalling_env(t *testing.T) {

	ai := NewEmptyConfig()

	err := ai.Unmarshal(fullConfig)
	require.NoError(t, err)

	res := env.MarshalEnvWithPrefix("MATRESHKA", &ai)

	expected := []env.EnvVal{
		{
			Name:  "MATRESHKA_APP_INFO_NAME",
			Value: "matreshka",
		},
		{
			Name:  "MATRESHKA_APP_INFO_VERSION",
			Value: "v0.0.1",
		},
		{
			Name:  "MATRESHKA_APP_INFO_STARTUP_DURATION",
			Value: time.Second * 10,
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_POSTGRES_HOST",
			Value: "localhost",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_POSTGRES_PORT",
			Value: uint64(5432),
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_POSTGRES_USER",
			Value: "matreshka",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_POSTGRES_PWD",
			Value: "matreshka",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_POSTGRES_DB_NAME",
			Value: "matreshka",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_POSTGRES_SSL_MODE",
			Value: "disable",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_REDIS_HOST",
			Value: "localhost",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_REDIS_PORT",
			Value: uint16(6379),
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_REDIS_USER",
			Value: "",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_REDIS_PWD",
			Value: "",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_REDIS_DB",
			Value: 0,
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_TELEGRAM_API_KEY",
			Value: "some_api_key",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_GRPC-RSCLI-EXAMPLE_CONNECTION_STRING",
			Value: "0.0.0.0:50051",
		},
		{
			Name:  "MATRESHKA_DATA_SOURCES_GRPC-RSCLI-EXAMPLE_MODULE",
			Value: "github.com/Red-Sock/rscli_example",
		},
		{
			Name:  "MATRESHKA_SERVERS_REST_PORT",
			Value: uint16(8080),
		},
		{
			Name:  "MATRESHKA_SERVERS_GRPC_PORT",
			Value: uint16(50051),
		},
	}

	require.ElementsMatch(t, expected, res)
}

func Test_unmarshal_env(t *testing.T) {
	const fileIn = `MATRESHKA_APP_INFO_NAME=matreshka
MATRESHKA_APP_INFO_VERSION=v0.0.1
MATRESHKA_APP_INFO_STARTUP_DURATION=10s
MATRESHKA_DATA_SOURCES_POSTGRES_HOST=localhost
MATRESHKA_DATA_SOURCES_POSTGRES_PORT=5432
MATRESHKA_DATA_SOURCES_POSTGRES_USER=matreshka
MATRESHKA_DATA_SOURCES_POSTGRES_PWD=matreshka
MATRESHKA_DATA_SOURCES_POSTGRES_DB_NAME=matreshka
MATRESHKA_DATA_SOURCES_POSTGRES_SSL_MODE=disable
MATRESHKA_DATA_SOURCES_REDIS_HOST=localhost
MATRESHKA_DATA_SOURCES_REDIS_PORT=6379
MATRESHKA_DATA_SOURCES_REDIS_USER=redis_matreshka
MATRESHKA_DATA_SOURCES_REDIS_PWD=redis_matreshka_pwd
MATRESHKA_DATA_SOURCES_REDIS_DB=2
MATRESHKA_DATA_SOURCES_TELEGRAM_API_KEY=some_api_key
MATRESHKA_DATA_SOURCES_GRPC-RSCLI-EXAMPLE_CONNECTION_STRING=0.0.0.0:50051
MATRESHKA_DATA_SOURCES_GRPC-RSCLI-EXAMPLE_MODULE=github.com/Red-Sock/rscli_example
MATRESHKA_SERVERS_REST_PORT=8080
MATRESHKA_SERVERS_GRPC_PORT=50051
`
	var c AppConfig
	env.UnmarshalWithPrefix("MATRESHKA", []byte(fileIn), &c)
	expected := AppConfig{
		AppInfo: AppInfo{
			Name:            "matreshka",
			Version:         "v0.0.1",
			StartupDuration: time.Second * 10,
		},
		DataSources: DataSources{
			&resources.Postgres{
				Name:    "postgres",
				Host:    "localhost",
				Port:    5432,
				User:    "matreshka",
				Pwd:     "matreshka",
				DbName:  "matreshka",
				SslMode: "disable",
			},
			&resources.Redis{
				Name: "redis",
				Host: "localhost",
				Port: 6379,
				User: "redis_matreshka",
				Pwd:  "redis_matreshka_pwd",
				Db:   2,
			},
			&resources.Telegram{
				Name:   "telegram",
				ApiKey: "some_api_key",
			},
			&resources.GRPC{
				Name:             "grpc_rscli_example",
				ConnectionString: "0.0.0.0:50051",
				Module:           "github.com/Red-Sock/rscli_example",
			},
		},
		Servers:     nil,
		Environment: nil,
	}
	require.Equal(t, c, expected)
}
