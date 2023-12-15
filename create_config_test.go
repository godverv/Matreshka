package matreshka

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/resources"
)

var (
	//go:embed tests/empty_config.yaml
	emptyConfig []byte
	//go:embed tests/app_config.yaml
	appConfig []byte
	//go:embed tests/resourced_config.yaml
	resourcedConfig []byte
	//go:embed tests/api_config.yaml
	apiConfig []byte
	//go:embed tests/api_half_empty_config.yaml
	apiHalfEmptyConfig []byte
	//go:embed tests/full_config.yaml
	fullConfig []byte
	//go:embed tests/environment_config.yaml
	environmentConfig []byte
	//go:embed tests/invalid_environment_config.yaml
	invalidEnvironmentConfig []byte
)

func Test_CreateEmptyConfig(t *testing.T) {
	t.Parallel()

	cfg := NewEmptyConfig()

	bytes, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, string(emptyConfig), string(bytes))
}

func Test_CreateConfigWithResources(t *testing.T) {
	t.Parallel()

	cfg := NewEmptyConfig()
	cfg.Name = "matreshka"
	cfg.Resources = append(cfg.Resources,
		&resources.Postgres{
			Name:    "postgres",
			Host:    "localhost",
			Port:    5432,
			DbName:  "matreshka",
			User:    "matreshka",
			Pwd:     "matreshka",
			SSLMode: "false",
		},
		&resources.Redis{
			Name: "redis",
			Host: "localhost",
			Port: 6379,
			User: "",
			Pwd:  "",
			Db:   0,
		},
		&resources.GRPC{
			Name:             "grpc_rscli_example",
			ConnectionString: "0.0.0.0:50051",
			Module:           "github.com/Red-Sock/rscli_example",
		},
		&resources.Telegram{
			Name:   "telegram",
			ApiKey: "some_api_key",
		},
	)

	bytes, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, string(resourcedConfig), string(bytes))
}

func Test_CreateConfigWithServers(t *testing.T) {
	t.Parallel()

	cfg := NewEmptyConfig()
	cfg.Name = "matreshka"
	cfg.Servers = append(cfg.Servers,
		&api.Rest{
			Name: "rest_server",
			Port: 8080,
		},
		&api.GRPC{
			Name:        "grpc_server",
			Port:        50051,
			GatewayPort: 50052,
		})

	apiMarshalled, err := cfg.Marshal()
	require.NoError(t, err)
	require.Equal(t, string(apiConfig), string(apiMarshalled))
}
