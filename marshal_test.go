package matreshka

import (
	"testing"
	"time"

	"github.com/Red-Sock/evon"
	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/resources"
	"github.com/godverv/matreshka/servers"
)

func Test_Marshal(t *testing.T) {
	t.Parallel()

	t.Run("YAML", func(t *testing.T) {
		t.Run("EMPTY", func(t *testing.T) {
			t.Parallel()
			cfg := NewEmptyConfig()

			bytes, err := cfg.Marshal()
			require.NoError(t, err)
			require.Equal(t, emptyConfig, bytes)
		})
		t.Run("WITH_RESOURCES", func(t *testing.T) {
			t.Parallel()

			cfg := NewEmptyConfig()
			cfg.Name = "matreshka"
			cfg.DataSources = append(cfg.DataSources,
				getPostgresClientTest(),
				getRedisClientTest(),
				getGRPCClientTest(),
				getTelegramClientTest(),
			)

			bytes, err := cfg.Marshal()
			require.NoError(t, err)
			require.Equal(t, resourcedConfig, bytes)
		})
		t.Run("WITH_SERVERS", func(t *testing.T) {
			t.Parallel()

			cfg := NewEmptyConfig()
			cfg.Name = "matreshka"
			cfg.Servers = append(cfg.Servers,
				getRestServerTest(),
				getGRPCServerTest(),
			)

			apiMarshalled, err := cfg.Marshal()
			require.NoError(t, err)
			require.Equal(t, string(apiConfig), string(apiMarshalled))
		})
	})

	t.Run("ENV_FULL", func(t *testing.T) {
		t.Parallel()

		ai := NewEmptyConfig()
		err := ai.Unmarshal(fullConfig)
		require.NoError(t, err)

		res := evon.MarshalEnvWithPrefix("MATRESHKA", &ai)

		expected := []evon.Node{
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
			{
				Name:  "MATRESHKA_ENVIRONMENT_STRING",
				Value: "not so basic 🤡 string",
			},
			{
				Name:  "MATRESHKA_ENVIRONMENT_INT",
				Value: 1,
			},
			{
				Name:  "MATRESHKA_ENVIRONMENT_DURATION",
				Value: "10s",
			},
			{
				Name:  "MATRESHKA_ENVIRONMENT_BOOL",
				Value: true,
			},
		}

		require.ElementsMatch(t, expected, res)
	})
}

func Test_Unmarshal(t *testing.T) {
	t.Parallel()
	t.Run("ENV_FULL", func(t *testing.T) {
		t.Parallel()
		c := NewEmptyConfig()
		evon.UnmarshalWithPrefix("MATRESHKA", dotEnvFullConfig, &c)
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
			Servers: Servers{
				&servers.Rest{
					Name: "rest",
					Port: 8080,
				},
				&servers.GRPC{
					Name: "grpc",
					Port: 50051,
				},
			},
			//Environment: Environment{
			//	"bool":     "true",
			//	"duration": "10s",
			//	"int":      "1",
			//	"string":   "not so basic 🤡 string",
			//},
		}
		require.Equal(t, c, expected)
	})
}
