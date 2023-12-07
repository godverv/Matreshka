package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/resources"
)

func Test_GetResources(t *testing.T) {
	t.Parallel()

	t.Run("OK_RESOURCES", func(t *testing.T) {
		cfg, err := ParseConfig(resourcedConfig)
		require.NoError(t, err)

		postgresCfg, err := cfg.Resources.Postgres("postgres")
		require.NoError(t, err)
		require.Equal(t, postgresCfg, &resources.Postgres{
			Name:    "postgres",
			Host:    "localhost",
			Port:    5432,
			User:    "matreshka",
			Pwd:     "matreshka",
			DbName:  "matreshka",
			SSLMode: "false",
		})

		redisCfg, err := cfg.Resources.Redis("redis")
		require.NoError(t, err)
		require.Equal(t, redisCfg, &resources.Redis{
			Name: "redis",
			Host: "localhost",
			Port: 6379,
			User: "",
			Pwd:  "",
			Db:   0,
		})

		grpcCfg, err := cfg.Resources.GRPC("grpc_rscli_example")
		require.NoError(t, err)
		require.Equal(t, grpcCfg, &resources.GRPC{
			Name:             "grpc_rscli_example",
			ConnectionString: "0.0.0.0:50051",
			Module:           "github.com/Red-Sock/rscli_example",
		})

		tgCfg, err := cfg.Resources.Telegram("telegram")
		require.NoError(t, err)
		require.Equal(t, tgCfg, &resources.Telegram{
			Name:   "telegram",
			ApiKey: "some_api_key",
		})
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
