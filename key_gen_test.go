package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_KeyGen(t *testing.T) {
	cfg := NewEmptyConfig()
	err := cfg.Unmarshal(fullConfig)
	require.NoError(t, err)

	keys := GenerateKeys(cfg)
	expected := []string{
		"rest",
		"grpc",
		"postgres",
		"redis",
		"telegram",
		"grpc_rscli_example",
		"int",
		"string",
		"bool",
		"duration",
	}
	require.ElementsMatch(t, keys, expected)
}
