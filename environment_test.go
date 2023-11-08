package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Environment(t *testing.T) {
	cfg, err := ParseConfig([]byte(environmentConfig))
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

func Test_InvalidEnvironment(t *testing.T) {
	cfg, err := ParseConfig([]byte(invalidEnvironmentConfig))
	require.NoError(t, err)

	t.Run("int", func(t *testing.T) {
		val, ok := cfg.TryGetInt("int")
		require.False(t, ok)
		require.Empty(t, val)
	})

	t.Run("string", func(t *testing.T) {
		val, ok := cfg.TryGetString("string")
		require.False(t, ok)
		require.Empty(t, val)
	})

	t.Run("bool", func(t *testing.T) {
		val, ok := cfg.TryGetBool("bool")
		require.False(t, ok)
		require.Empty(t, val)
	})

	t.Run("duration", func(t *testing.T) {
		val, ok := cfg.TryGetDuration("duration")
		require.False(t, ok)
		require.Empty(t, val)
	})
}
