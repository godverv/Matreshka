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

	expected := ApplicationKeys{
		Servers:     []string{"rest", "grpc"},
		DataSources: []string{"postgres", "redis", "telegram", "grpc_rscli_example"},
	}
	require.Equal(t, keys, expected)
}
