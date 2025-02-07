package matreshka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.vervstack.ru/matreshka/environment"
	config "go.vervstack.ru/matreshka/internal/config_test"
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
			environment.MustNewVariable("new_unknown", "nil"),
		})

		customEnvConf := &config.EnvironmentConfig{}

		err := env.ParseToStruct(customEnvConf)
		require.ErrorIs(t, err, ErrNotFound)

		expected := &config.EnvironmentConfig{}
		require.Equal(t, expected, customEnvConf)
	})

	t.Run("MARSHAL", func(t *testing.T) {
		ac := AppConfig{
			Environment: getEnvironmentVariables(),
		}
		// TODO RSI-294: add tests after reformating the way yamls are created
		bytes, err := ac.Marshal()
		require.NoError(t, err)

		_ = bytes
	})
}
