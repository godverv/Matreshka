package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetApi(t *testing.T) {
	t.Parallel()

	t.Run("OK_API", func(t *testing.T) {
		cfg, err := ParseConfig(apiConfig)
		require.NoError(t, err)

		grpcCfg, err := cfg.Servers.GRPC("grpc_server")
		require.NoError(t, err)
		require.Equal(t, grpcCfg, getGRPCServerTest())

		restCfg, err := cfg.Servers.REST("rest_server")
		require.NoError(t, err)
		require.Equal(t, restCfg, getRestServerTest())
	})

	t.Run("ERROR_NOT_FOUND_API", func(t *testing.T) {
		cfg, err := ParseConfig(emptyConfig)
		require.NoError(t, err)

		grpcCfg, err := cfg.Servers.GRPC("jrpc_server")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, grpcCfg)

		restCfg, err := cfg.Servers.REST("full_rest_server")
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, restCfg)
	})

	t.Run("ERROR_INVALID_TYPE_API", func(t *testing.T) {
		cfg, err := ParseConfig(apiConfig)
		require.NoError(t, err)

		grpcCfg, err := cfg.Servers.GRPC("rest_server")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, grpcCfg)

		restCfg, err := cfg.Servers.REST("grpc_server")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Nil(t, restCfg)
	})

	t.Run("OK_HALF_EMPTY_API_CONFIG", func(t *testing.T) {
		_, err := ParseConfig(apiHalfEmptyConfig)
		require.NoError(t, err)
	})
}
