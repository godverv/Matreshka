package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/internal/env_parser"
)

func Test_GenerateGoConfigKeys(t *testing.T) {
	c, err := ParseConfig([]byte(fullConfig))
	require.NoError(t, err)

	expected := []env_parser.EnvVal{
		{
			Name:  "matreshka_bool",
			Value: true,
		},
		{
			Name:  "matreshka_duration",
			Value: "10s",
		},
		{
			Name:  "matreshka_int",
			Value: 1,
		},
		{
			Name:  "matreshka_string",
			Value: "not so basic 🤡 string",
		},
	}

	res, err := GenerateKeys(*c)
	require.NoError(t, err)
	require.Equal(t, expected, res)
}
