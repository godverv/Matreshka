package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetResources(t *testing.T) {
	t.Parallel()

	t.Run("OK_RESOURCES", func(t *testing.T) {
		cfg, err := ParseConfig(resourcedConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.Resources.Postgres("postgres")
		require.NoError(t, err)
		require.Equal(t, postgresCfg, getPostgresClientTest())

		redisCfg, err := cfg.Resources.Redis("redis")
		require.NoError(t, err)
		require.Equal(t, redisCfg, getRedisClientTest())

		grpcCfg, err := cfg.Resources.GRPC("grpc_rscli_example")
		require.NoError(t, err)
		require.Equal(t, grpcCfg, getGRPCClientTest())

		tgCfg, err := cfg.Resources.Telegram("telegram")
		require.NoError(t, err)
		require.Equal(t, tgCfg, getTelegramClientTest())
	})

	t.Run("ERROR_NOT_FOUND_RESOURCE", func(t *testing.T) {
		cfg, err := ParseConfig(emptyConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.Resources.Postgres("postgres")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, postgresCfg)

		redisCfg, err := cfg.Resources.Redis("redis")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, redisCfg)

		grpcCfg, err := cfg.Resources.GRPC("grpc_rscli_example")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, grpcCfg)

		tgCfg, err := cfg.Resources.Telegram("redis")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, tgCfg)
	})

	t.Run("ERROR_INVALID_TYPE_API", func(t *testing.T) {
		cfg, err := ParseConfig(resourcedConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.Resources.Redis("postgres")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, postgresCfg)

		redisCfg, err := cfg.Resources.GRPC("redis")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, redisCfg)

		grpcCfg, err := cfg.Resources.Postgres("grpc_rscli_example")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, grpcCfg)

		tgCfg, err := cfg.Resources.Telegram("redis")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, tgCfg)
	})
}
