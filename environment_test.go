package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.verv.tech/matreshka/environment"
	config "go.verv.tech/matreshka/internal/config_test"
)

func Test_Environment(t *testing.T) {
	t.Parallel()

	t.Run("PARSE_ENV_TO_STRUCT", func(t *testing.T) {
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

	t.Run("PARSE_ENV_MORE_THAN_HAVE_IN_STRUCT", func(t *testing.T) {
		t.Parallel()

		env := Environment([]*environment.Variable{
			{
				Name:  "new_unknown",
				Type:  environment.VariableTypeStr,
				Value: "nil",
			},
		})

		customEnvConf := &config.EnvironmentConfig{}

		err := env.ParseToStruct(customEnvConf)
		require.ErrorIs(t, err, ErrNotFound)

		expected := &config.EnvironmentConfig{}
		require.Equal(t, expected, customEnvConf)
	})
}
