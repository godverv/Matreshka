package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	config "github.com/godverv/matreshka/config_test"
)

func Test_Environment(t *testing.T) {
	t.Parallel()

	t.Run("gen_go_struct", func(t *testing.T) {
		t.Parallel()

		env := Environment(getEnvironmentVariables())

		generatedCustomGoStruct := env.GenerateCustomGoStruct()
		require.Equal(t, string(goCustomEnvStruct), string(generatedCustomGoStruct))
	})

	t.Run("parse_env_to_struct", func(t *testing.T) {
		t.Parallel()

		env := Environment(getEnvironmentVariables())

		customEnvConf := &config.EnvironmentConfig{}

		err := env.ParseToStruct(customEnvConf)
		require.NoError(t, err)

		expected := &config.EnvironmentConfig{
			AvailablePorts:                   []int{10, 12, 34, 35, 36, 37, 38, 39, 40},
			CreditPercent:                    0.01,
			CreditPercentsBasedOnYearOfBirth: []float64{0.01, 0.02, 0.03, 0.04},
			DatabaseMaxConnections:           1,
			OneOfWelcomeString:               "one",
			RequestTimeout:                   time.Second * 10,
			TrueFalser:                       true,
			UsernamesToBan:                   []string{"hacker228", "mothe4acker"},
			WelcomeString:                    "not so basic 🤡 string",
		}
		require.Equal(t, expected, customEnvConf)
	})
}
