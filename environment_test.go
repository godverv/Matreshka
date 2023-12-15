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
			cfg.GetInt("matreshka_int"))
	})

	t.Run("string", func(t *testing.T) {
		require.Equal(t,
			"not so basic 🤡 string",
			cfg.GetString("matreshka_string"))
	})

	t.Run("bool", func(t *testing.T) {
		require.Equal(t,
			true,
			cfg.GetBool("matreshka_bool"))
	})

	t.Run("duration", func(t *testing.T) {
		require.Equal(t,
			time.Second*10,
			cfg.GetDuration("matreshka_duration"))
	})

	t.Run("slice", func(t *testing.T) {
		var s []any
		err = ReadSliceFromConfig(cfg, "matreshka_slice", &s)
		require.NoError(t, err)
		require.Equal(t, []any{"1", "2", "3", "4"}, s)
	})
}

func Test_EnvironmentInvalid(t *testing.T) {
	t.Parallel()

	cfg, err := ParseConfig(invalidEnvironmentConfig)
	require.NoError(t, err)

	t.Run("int", func(t *testing.T) {
		val, err := cfg.TryGetInt("matreshka_int")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Empty(t, val)
	})

	t.Run("string", func(t *testing.T) {
		val, err := cfg.TryGetString("matreshka_string")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Empty(t, val)
	})

	t.Run("bool", func(t *testing.T) {
		val, err := cfg.TryGetBool("matreshka_bool")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Empty(t, val)
	})

	t.Run("duration", func(t *testing.T) {
		val, err := cfg.TryGetDuration("matreshka_duration")
		require.ErrorIs(t, err, ErrUnexpectedType)
		require.Empty(t, val)
	})
}

func Test_EnvironmentNotFound(t *testing.T) {
	t.Parallel()

	cfg, err := ParseConfig(emptyConfig)
	require.NoError(t, err)

	t.Run("int", func(t *testing.T) {
		val, err := cfg.TryGetInt("matreshka_not_found_int")
		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, val)
	})

	t.Run("string", func(t *testing.T) {
		val, err := cfg.TryGetString("matreshka_string")
		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, val)
	})

	t.Run("bool", func(t *testing.T) {
		val, err := cfg.TryGetBool("matreshka_bool")
		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, val)
	})

	t.Run("duration", func(t *testing.T) {
		val, err := cfg.TryGetDuration("matreshka_duration")
		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, val)
	})
}
