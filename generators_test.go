package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenerateGoConfigKeys(t *testing.T) {
	c, err := ParseConfig([]byte(fullConfig))
	require.NoError(t, err)

	expectedKeys := []string{
		"matreshka_bool",
		"matreshka_duration",
		"matreshka_int",
		"matreshka_string",
	}
	expectedValues := []any{
		1,
		true,
		"not so basic 🤡 string",
		"10s",
	}

	keys, values, err := GenerateEnvironmentKeys(*c)
	require.NoError(t, err)
	require.Equal(t, expectedKeys, keys)
	require.Equal(t, expectedValues, values)
}
