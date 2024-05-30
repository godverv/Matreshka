package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetApi(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		cfg, err := ParseConfig(apiConfig)
		require.NoError(t, err)

		grpcCfg, err := cfg.Servers.GRPC("grpc")
		require.NoError(t, err)
		require.Equal(t, grpcCfg, getGRPCServerTest())

		restCfg, err := cfg.Servers.REST("rest")
		require.NoError(t, err)
		require.Equal(t, restCfg, getRestServerTest())
	})

	t.Run("ERROR_NOT_FOUND", func(t *testing.T) {
		cfg, err := ParseConfig(emptyConfig)
		require.NoError(t, err)

		grpcCfg, err := cfg.Servers.GRPC("jrpc_server")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, grpcCfg)

		restCfg, err := cfg.Servers.REST("full_rest_server")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, restCfg)
	})

	t.Run("ERROR_INVALID_TYPE", func(t *testing.T) {
		cfg, err := ParseConfig(apiConfig)
		require.NoError(t, err)

		grpcCfg, err := cfg.Servers.GRPC("rest")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, grpcCfg)

		restCfg, err := cfg.Servers.REST("grpc")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, restCfg)
	})

	t.Run("OK_HALF_EMPTY", func(t *testing.T) {
		_, err := ParseConfig(apiHalfEmptyConfig)
		require.NoError(t, err)
	})
}
