package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetResources(t *testing.T) {
	t.Parallel()

	t.Run("OK_RESOURCES_WITH_PREFIX", func(t *testing.T) {
		cfg, err := ParseConfig(resourcedConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.DataSources.Postgres(resourcePrefix + "postgres")
		require.NoError(t, err)
		require.Equal(t, postgresCfg, getPostgresClientTest())

		redisCfg, err := cfg.DataSources.Redis(resourcePrefix + "redis")
		require.NoError(t, err)
		require.Equal(t, redisCfg, getRedisClientTest())

		grpcCfg, err := cfg.DataSources.GRPC(resourcePrefix + "grpc_rscli_example")
		require.NoError(t, err)
		require.Equal(t, grpcCfg, getGRPCClientTest())

		tgCfg, err := cfg.DataSources.Telegram(resourcePrefix + "telegram")
		require.NoError(t, err)
		require.Equal(t, tgCfg, getTelegramClientTest())
	})

	t.Run("OK_RESOURCES_WITHOUT_PREFIX", func(t *testing.T) {
		cfg, err := ParseConfig(resourcedConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.DataSources.Postgres("postgres")
		require.NoError(t, err)
		require.Equal(t, postgresCfg, getPostgresClientTest())

		redisCfg, err := cfg.DataSources.Redis("redis")
		require.NoError(t, err)
		require.Equal(t, redisCfg, getRedisClientTest())

		grpcCfg, err := cfg.DataSources.GRPC("grpc_rscli_example")
		require.NoError(t, err)
		require.Equal(t, grpcCfg, getGRPCClientTest())

		tgCfg, err := cfg.DataSources.Telegram("telegram")
		require.NoError(t, err)
		require.Equal(t, tgCfg, getTelegramClientTest())
	})

	t.Run("ERROR_NOT_FOUND_RESOURCE", func(t *testing.T) {
		cfg, err := ParseConfig(emptyConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.DataSources.Postgres(resourcePrefix + "postgres")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, postgresCfg)

		redisCfg, err := cfg.DataSources.Redis(resourcePrefix + "redis")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, redisCfg)

		grpcCfg, err := cfg.DataSources.GRPC(resourcePrefix + "grpc_rscli_example")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, grpcCfg)

		tgCfg, err := cfg.DataSources.Telegram(resourcePrefix + "redis")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, tgCfg)
	})

	t.Run("ERROR_INVALID_TYPE_API", func(t *testing.T) {
		cfg, err := ParseConfig(resourcedConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.DataSources.Redis(resourcePrefix + "postgres")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, postgresCfg)

		redisCfg, err := cfg.DataSources.GRPC(resourcePrefix + "redis")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, redisCfg)

		grpcCfg, err := cfg.DataSources.Postgres(resourcePrefix + "grpc_rscli_example")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, grpcCfg)

		tgCfg, err := cfg.DataSources.Telegram(resourcePrefix + "redis")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, tgCfg)
	})
}
