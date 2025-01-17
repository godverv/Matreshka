package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.redsock.ru/evon"
)

const (
	postgresResourceName = "postgres"
	redisResourceName    = "redis"
	grpcResourceName     = "grpc_rscli_example"
	grpcResourceModule   = "github.com/Red-Sock/rscli_example"
	telegramResourceName = "telegram"
)

func Test_GetResources(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		cfg, err := ParseConfig(fullConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.DataSources.Postgres(postgresResourceName)
		require.NoError(t, err)
		require.Equal(t, postgresCfg, getPostgresClientTest())

		redisCfg, err := cfg.DataSources.Redis(redisResourceName)
		require.NoError(t, err)
		require.Equal(t, redisCfg, getRedisClientTest())

		grpcCfg, err := cfg.DataSources.GRPC(grpcResourceName)
		require.NoError(t, err)
		require.Equal(t, grpcCfg, getGRPCClientTest())

		tgCfg, err := cfg.DataSources.Telegram(telegramResourceName)
		require.NoError(t, err)
		require.Equal(t, tgCfg, getTelegramClientTest())
	})

	t.Run("ERROR_RESOURCE_NOT_FOUND", func(t *testing.T) {
		cfg, err := ParseConfig(emptyConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.DataSources.Postgres("postgres")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, postgresCfg)

		redisCfg, err := cfg.DataSources.Redis("redis")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, redisCfg)

		grpcCfg, err := cfg.DataSources.GRPC("grpc_rscli_example")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, grpcCfg)

		tgCfg, err := cfg.DataSources.Telegram("redis")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, tgCfg)
	})

	t.Run("ERROR_INVALID_RESOURCE_TYPE", func(t *testing.T) {
		cfg, err := ParseConfig(fullConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.DataSources.Redis("postgres")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, postgresCfg)

		redisCfg, err := cfg.DataSources.GRPC("redis")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, redisCfg)

		grpcCfg, err := cfg.DataSources.Postgres("grpc_rscli_example")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, grpcCfg)

		tgCfg, err := cfg.DataSources.Telegram("redis")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, tgCfg)
	})
}

func Test_DataSources(t *testing.T) {
	t.Parallel()

	t.Run("YAML", func(t *testing.T) {
		t.Run("Marshal", func(t *testing.T) {
			cfg := AppConfig{
				DataSources: getResourcesFull(),
			}

			marshalled, err := cfg.Marshal()
			require.NoError(t, err)

			require.Equal(t, string(marshalled), string(fullResourcesConfig))
		})
		t.Run("Unmarshal", func(t *testing.T) {
			actualConfig := &AppConfig{}

			err := actualConfig.Unmarshal(fullResourcesConfig)
			require.NoError(t, err)
			expectedConfig := &AppConfig{
				DataSources: getResourcesFull(),
				Servers:     Servers{},
				Environment: Environment{},
			}

			require.Equal(t, expectedConfig, actualConfig)
		})
	})

	t.Run("EVON", func(t *testing.T) {
		t.Run("Marshal", func(t *testing.T) {
			cfg := AppConfig{
				DataSources: getResourcesFull(),
			}

			marshalledEvonNodes, err := evon.MarshalEnv(cfg)
			require.NoError(t, err)
			evonMarshalled := evon.Marshal(marshalledEvonNodes.InnerNodes)
			require.Equal(t, string(evonMarshalled), string(fullResourcesEnvConfig))
		})

		t.Run("Unmarshal", func(t *testing.T) {
			actualConfig := AppConfig{}
			err := evon.Unmarshal(fullResourcesEnvConfig, &actualConfig)
			require.NoError(t, err)

			expectedConfig := AppConfig{
				DataSources: getResourcesFull(),
			}

			require.Equal(t, expectedConfig, actualConfig)
		})
	})
}
