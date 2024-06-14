package matreshka

import (
	"testing"
	"time"

	"github.com/Red-Sock/evon"
	"github.com/stretchr/testify/require"
)

func Test_Marshal(t *testing.T) {
	t.Parallel()

	t.Run("ENV", func(t *testing.T) {
		t.Parallel()

		srcConfig := NewEmptyConfig()
		err := srcConfig.Unmarshal(fullConfig)
		require.NoError(t, err)

		evonNodes, err := evon.MarshalEnvWithPrefix("MATRESHKA", &srcConfig)
		require.NoError(t, err)

		expected := getEvonFullConfig()
		require.Equal(t, evonNodes, expected)

		evonBytes := evon.Marshal(evonNodes.InnerNodes)
		require.Equal(t, string(dotEnvFullConfig), string(evonBytes),
			"expected to match environment in dotenv file")
	})
}

func Test_Unmarshal(t *testing.T) {
	t.Parallel()

	fullConfigExpected := AppConfig{
		AppInfo: AppInfo{
			Name:            "matreshka",
			Version:         "v0.0.1",
			StartupDuration: time.Second * 10,
		},
		DataSources: DataSources{
			getPostgresClientTest(),
			getRedisClientTest(),
			getTelegramClientTest(),
			getGRPCClientTest(),
		},
		Servers: Servers{
			getRestServerTest(),
			getGRPCServerTest(),
		},
		Environment: getEnvironmentVariables(),
	}
	t.Run("ENV", func(t *testing.T) {
		t.Parallel()

		unmarshalledConfig := NewEmptyConfig()
		evon.UnmarshalWithPrefix("MATRESHKA", dotEnvFullConfig, &unmarshalledConfig)

		require.Equal(t, unmarshalledConfig, fullConfigExpected)
	})

	t.Run("YAML", func(t *testing.T) {
		t.Parallel()

		unmarshalledConfig := NewEmptyConfig()
		err := unmarshalledConfig.Unmarshal(fullConfig)
		require.NoError(t, err)

		require.Equal(t, unmarshalledConfig, fullConfigExpected)
	})
}
