package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/resources"
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

func Test_ResourceObfuscation(t *testing.T) {
	t.Parallel()

	cfg, err := ParseConfig(fullConfig)
	require.NoError(t, err)

	postgresCfg, _ := cfg.DataSources.Postgres(postgresResourceName)
	expectedPg := &resources.Postgres{
		Name:   postgresResourceName,
		Host:   "localhost",
		Port:   5432,
		User:   "postgres",
		Pwd:    "postgres",
		DbName: "master",
	}
	require.Equal(t, expectedPg, postgresCfg.Obfuscate())

	redisCfg, _ := cfg.DataSources.Redis(redisResourceName)
	expectedRedis := &resources.Redis{
		Name: redisResourceName,
		Host: "localhost",
		Port: 6379,
	}

	require.Equal(t, expectedRedis, redisCfg.Obfuscate())

	grpcCfg, _ := cfg.DataSources.GRPC(grpcResourceName)
	expectedGrpc := &resources.GRPC{
		Name:             grpcResourceName,
		ConnectionString: "localhost:50051",
		Module:           grpcResourceModule,
	}
	require.Equal(t, expectedGrpc, grpcCfg.Obfuscate())

	tgCfg, _ := cfg.DataSources.Telegram(telegramResourceName)
	expectedTelegram := &resources.Telegram{
		Name: telegramResourceName,
	}
	require.Equal(t, expectedTelegram, tgCfg.Obfuscate())
}
