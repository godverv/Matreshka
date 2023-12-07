package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_EnvironmentOk(t *testing.T) {
	t.Parallel()

	cfg, err := ParseConfig(environmentConfig)
	require.NoError(t, err)

	t.Run("int", func(t *testing.T) {
		require.Equal(t,
			1,
			cfg.GetInt("int"))
	})

	t.Run("string", func(t *testing.T) {
		require.Equal(t,
			"not so basic 🤡 string",
			cfg.GetString("string"))
	})

	t.Run("bool", func(t *testing.T) {
		require.Equal(t,
			true,
			cfg.GetBool("bool"))
	})

	t.Run("duration", func(t *testing.T) {
		require.Equal(t,
			time.Second*10,
			cfg.GetDuration("duration"))
	})
}

func Test_EnvironmentInvalid(t *testing.T) {
	t.Parallel()

	cfg, err := ParseConfig(invalidEnvironmentConfig)
	require.NoError(t, err)

	t.Run("int", func(t *testing.T) {
		val, err := cfg.TryGetInt("int")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Empty(t, val)
	})

	t.Run("string", func(t *testing.T) {
		val, err := cfg.TryGetString("string")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Empty(t, val)
	})

	t.Run("bool", func(t *testing.T) {
		val, err := cfg.TryGetBool("bool")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Empty(t, val)
	})

	t.Run("duration", func(t *testing.T) {
		val, err := cfg.TryGetDuration("duration")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Empty(t, val)
	})
}

func Test_EnvironmentNotFound(t *testing.T) {
	t.Parallel()

	cfg, err := ParseConfig(emptyConfig)
	require.NoError(t, err)

	t.Run("int", func(t *testing.T) {
		val, err := cfg.TryGetInt("not_found_int")
		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, val)
	})

	t.Run("string", func(t *testing.T) {
		val, err := cfg.TryGetString("string")
		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, val)
	})

	t.Run("bool", func(t *testing.T) {
		val, err := cfg.TryGetBool("bool")
		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, val)
	})

	t.Run("duration", func(t *testing.T) {
		val, err := cfg.TryGetDuration("duration")
		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, val)
	})
}
