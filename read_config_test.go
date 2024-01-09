package matreshka

import (
	"context"
	"testing"
	"time"

	"github.com/godverv/matreshka-be/pkg/api/matreshka_api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/mocks"
	"github.com/godverv/matreshka/resources"
)

func Test_ReadAppConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(appConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.AppInfo = AppInfo{
		Name:            "matreshka",
		Version:         "0.0.1",
		StartupDuration: 10 * time.Second,
	}

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_ReadDataSourceConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(resourcedConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.Name = "matreshka"

	cfgExpect.Resources = append(cfgExpect.Resources,
		getPostgresClientTest(),
		getRedisClientTest(),
		getGRPCClientTest(),
		getTelegramClientTest(),
	)

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_ReadApiConfig(t *testing.T) {
	t.Parallel()

	cfgGot, err := ParseConfig(apiConfig)
	require.NoError(t, err)

	cfgExpect := NewEmptyConfig()
	cfgExpect.Name = "matreshka"
	cfgExpect.Servers = []api.Api{
		getRestServerTest(),
		getGRPCServerTest(),
	}

	require.Equal(t, cfgExpect, cfgGot)
}

func Test_ReadWithApi(t *testing.T) {
	backend := mocks.NewMatreshkaBeAPIClientMock(t)
	defer backend.MinimockFinish()

	backend.GetConfigRawMock.Set(
		func(_ context.Context, in *matreshka_api.GetConfigRaw_Request, opts ...grpc.CallOption) (gp1 *matreshka_api.GetConfigRaw_Response, err error) {
			return &matreshka_api.GetConfigRaw_Response{
				Config: string(fullConfig),
			}, nil
		})

	cfgGot := getViaApi(backend, "MATRESHKA_TEST")

	cfgExpect := NewEmptyConfig()
	cfgExpect.Name = "matreshka"
	cfgExpect.Version = "0.0.1"
	cfgExpect.StartupDuration = time.Second * 10
	cfgExpect.Resources = []resources.Resource{
		getPostgresClientTest(),
		getRedisClientTest(),
		getTelegramClientTest(),
		getGRPCClientTest(),
	}
	cfgExpect.Servers = []api.Api{
		getRestServerTest(),
		getGRPCServerTest(),
	}
	cfgExpect.Environment = map[string]interface{}{
		"matreshka_bool":     true,
		"matreshka_duration": "10s",
		"matreshka_int":      1,
		"matreshka_string":   "not so basic 🤡 string",
	}

	require.Equal(t, cfgExpect, cfgGot)
}
