package matreshka

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.redsock.ru/evon"

	"go.vervstack.ru/matreshka/service_discovery"
)

func Test_ServiceDiscovery(t *testing.T) {
	t.Parallel()

	t.Run("YAML", func(t *testing.T) {
		t.Run("Marshal_OK", func(t *testing.T) {
			t.Parallel()
			var cfgIn AppConfig
			cfgIn.ServiceDiscovery = getConfigServiceDiscovery()

			marshaled, err := cfgIn.Marshal()
			require.NoError(t, err)
			require.Equal(t, string(serviceDiscoveryConfig), string(marshaled))
		})

		t.Run("Unmarshal_OK", func(t *testing.T) {
			t.Parallel()
			actualConfig, err := ParseConfig(serviceDiscoveryConfig)
			require.NoError(t, err)

			expectedServiceDiscovery := getConfigServiceDiscovery()

			require.Equal(t, expectedServiceDiscovery, actualConfig.ServiceDiscovery)
		})
	})

	t.Run("ENV", func(t *testing.T) {
		t.Run("Marshal", func(t *testing.T) {
			t.Parallel()
			var cfgIn AppConfig
			cfgIn.ServiceDiscovery = getConfigServiceDiscovery()

			marshalledNodes, err := evon.MarshalEnvWithPrefix("MATRESHKA", cfgIn)
			require.NoError(t, err)
			marshalled := evon.Marshal(marshalledNodes.InnerNodes)
			require.Equal(t, string(serviceDiscoveryEnvConfig), string(marshalled))
		})
	})
}

func getConfigServiceDiscovery() ServiceDiscovery {
	return ServiceDiscovery{
		MakoshUrl:   "localhost:1281",
		MakoshToken: "1256",
		Overrides: service_discovery.Overrides{
			{
				ServiceName: "matreshka",
				Urls: []string{
					"localhost:1257",
				},
			},
		},
	}
}
