package matreshka

import (
	"testing"

	"github.com/Red-Sock/evon"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/server"
)

func Test_Servers(t *testing.T) {
	t.Run("YAML", func(t *testing.T) {
		t.Run("Marshal_OK", func(t *testing.T) {
			t.Parallel()

			var cfgIn AppConfig
			cfgIn.Servers = getConfigServersFull()

			marshaled, err := cfgIn.Marshal()
			require.NoError(t, err)

			var actual map[any]any

			require.NoError(t, yaml.Unmarshal(marshaled, &actual))

			expected := map[any]any{
				"servers": map[any]any{
					8080: map[string]any{
						"/{FS}": map[string]any{
							"dist": "web/dist",
						},
					},
					50051: map[string]any{
						"/{GRPC}": map[string]any{
							"module":  "pkg/matreshka_be_api",
							"gateway": "/api",
						},
					},
				},
			}

			require.Equal(t, expected, actual)
		})
		t.Run("Unmarshal_OK", func(t *testing.T) {
			t.Parallel()

			cfg, err := ParseConfig(apiConfig)
			require.NoError(t, err)

			servers := getConfigServersFull()
			require.Equal(t, cfg.Servers, servers)
		})
	})

	t.Run("ENV", func(t *testing.T) {
		t.Run("Marshal", func(t *testing.T) {
			t.Parallel()

			var cfgIn AppConfig
			cfgIn.Servers = getConfigServersFull()

			marshaledNodes, err := cfgIn.Servers.MarshalEnv("MATRESHKA_SERVERS")
			require.NoError(t, err)

			marshalledBytes := evon.Marshal(marshaledNodes)
			require.Equal(t, string(apiEnvConfig), string(marshalledBytes))
		})
		t.Run("Unmarshal", func(t *testing.T) {
			t.Parallel()
			cfg := NewEmptyConfig()
			err := evon.UnmarshalWithPrefix("MATRESHKA", apiEnvConfig, &cfg)
			require.NoError(t, err)

			servers := getConfigServersFull()
			require.Equal(t, cfg.Servers, servers)
		})
	})

}

func getConfigServersFull() Servers {
	return Servers{
		8080: {
			GRPC: map[string]*server.GRPC{},
			FS: map[string]*server.FS{
				"/{FS}": {
					Dist: "web/dist",
				},
			},
			HTTP: map[string]*server.HTTP{},
		},
		50051: {
			GRPC: map[string]*server.GRPC{
				"/{GRPC}": {
					Module:  "pkg/matreshka_be_api",
					Gateway: "/api",
				},
			},
			FS:   map[string]*server.FS{},
			HTTP: map[string]*server.HTTP{},
		},
	}
}
