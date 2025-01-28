package matreshka

import (
	"os"
	"path"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.verv.tech/matreshka/resources"
)

func Test_ReadConfig(t *testing.T) {
	tmpDirPath := path.Join(os.TempDir(), t.Name())
	require.NoError(t, os.MkdirAll(tmpDirPath, os.ModePerm))

	t.Run("OK_EMPTY", func(t *testing.T) {
		cfgPath := path.Join(tmpDirPath, path.Base(t.Name())+".yaml")
		defer func() {
			require.NoError(t, os.RemoveAll(cfgPath))
		}()

		require.NoError(t,
			os.WriteFile(
				cfgPath,
				emptyConfig,
				os.ModePerm))

		cfg, err := getFromFile(cfgPath)
		require.NoError(t, err)
		require.Equal(t, cfg, NewEmptyConfig())
	})

	t.Run("OK_READ_FULL_FROM_FILE", func(t *testing.T) {
		t.Parallel()

		cfgActual, err := ParseConfig(fullConfig)
		require.NoError(t, err)

		cfgExpect := getFullConfigTest()
		for _, s := range cfgExpect.Servers {
			s.Name = ""
		}

		require.Equal(t, cfgExpect.AppInfo, cfgActual.AppInfo)
		require.Equal(t, cfgExpect.Environment, cfgActual.Environment)
		require.Equal(t, cfgExpect.DataSources, cfgActual.DataSources)
		require.Equal(t, cfgExpect.ServiceDiscovery, cfgActual.ServiceDiscovery)
		require.Equal(t, cfgExpect.Servers, cfgActual.Servers)
	})

	t.Run("OK_READ_FULL_FROM_ENVIRONMENT_EVON_FORMAT", func(t *testing.T) {
		setupFullEnvConfigWithModuleName(t)

		cfgActual, err := ReadConfigs()
		require.NoError(t, err)

		cfgExpect := getFullConfigTest()

		sort.Slice(cfgExpect.DataSources, func(i, j int) bool {
			return cfgExpect.DataSources[i].GetName() > cfgExpect.DataSources[j].GetName()
		})

		sort.Slice(cfgActual.DataSources, func(i, j int) bool {
			return cfgActual.DataSources[i].GetName() > cfgActual.DataSources[j].GetName()
		})

		sort.Slice(cfgExpect.Environment, func(i, j int) bool {
			return cfgExpect.Environment[i].Name > cfgExpect.Environment[j].Name
		})

		sort.Slice(cfgActual.Environment, func(i, j int) bool {
			return cfgActual.Environment[i].Name > cfgActual.Environment[j].Name
		})

		require.Equal(t, cfgExpect, cfgActual)
	})

	t.Run("OK_READ_FULL_FROM_ENVIRONMENT_USING_MODULE_NAME_SHORT", func(t *testing.T) {
		cfgPath := path.Join(tmpDirPath, path.Base(t.Name())+".yaml")
		defer func() {
			require.NoError(t, os.RemoveAll(cfgPath))
		}()

		require.NoError(t,
			os.WriteFile(
				cfgPath,
				appInfoConfigShortName,
				os.ModePerm))

		setupFullEnvConfigWithModuleName(t)

		require.NoError(t, os.Setenv(VervName, ""))

		cfgActual, err := ReadConfigs(cfgPath)
		require.NoError(t, err)

		cfgExpect := getFullConfigTest()

		sort.Slice(cfgExpect.DataSources, func(i, j int) bool {
			return cfgExpect.DataSources[i].GetName() > cfgExpect.DataSources[j].GetName()
		})

		sort.Slice(cfgActual.DataSources, func(i, j int) bool {
			return cfgActual.DataSources[i].GetName() > cfgActual.DataSources[j].GetName()
		})

		sort.Slice(cfgExpect.Environment, func(i, j int) bool {
			return cfgExpect.Environment[i].Name > cfgExpect.Environment[j].Name
		})

		sort.Slice(cfgActual.Environment, func(i, j int) bool {
			return cfgActual.Environment[i].Name > cfgActual.Environment[j].Name
		})

		require.Equal(t, cfgExpect, cfgActual)
	})
	t.Run("OK_READ_FULL_FROM_ENVIRONMENT_USING_MODULE_NAME_FULL", func(t *testing.T) {
		cfgPath := path.Join(tmpDirPath, path.Base(t.Name())+".yaml")
		defer func() {
			require.NoError(t, os.RemoveAll(cfgPath))
		}()

		require.NoError(t,
			os.WriteFile(
				cfgPath,
				appInfoConfigFullName,
				os.ModePerm))

		setupFullEnvConfigWithModuleName(t)

		require.NoError(t, os.Setenv(VervName, ""))

		cfgActual, err := ReadConfigs(cfgPath)
		require.NoError(t, err)

		cfgExpect := getFullConfigTest()

		sort.Slice(cfgExpect.DataSources, func(i, j int) bool {
			return cfgExpect.DataSources[i].GetName() > cfgExpect.DataSources[j].GetName()
		})

		sort.Slice(cfgActual.DataSources, func(i, j int) bool {
			return cfgActual.DataSources[i].GetName() > cfgActual.DataSources[j].GetName()
		})

		sort.Slice(cfgExpect.Environment, func(i, j int) bool {
			return cfgExpect.Environment[i].Name > cfgExpect.Environment[j].Name
		})

		sort.Slice(cfgActual.Environment, func(i, j int) bool {
			return cfgActual.Environment[i].Name > cfgActual.Environment[j].Name
		})

		require.Equal(t, cfgExpect, cfgActual)
	})
	t.Run("OK_READ_FULL_FROM_ENVIRONMENT_WITHOUT_MODULE_NAME", func(t *testing.T) {
		setupFullEnvConfigWithoutModuleName(t)

		require.NoError(t, os.Setenv(VervName, ""))

		cfgActual, err := ReadConfigs()
		require.NoError(t, err)

		cfgExpect := getFullConfigTest()

		sort.Slice(cfgExpect.DataSources, func(i, j int) bool {
			return cfgExpect.DataSources[i].GetName() > cfgExpect.DataSources[j].GetName()
		})

		sort.Slice(cfgActual.DataSources, func(i, j int) bool {
			return cfgActual.DataSources[i].GetName() > cfgActual.DataSources[j].GetName()
		})

		sort.Slice(cfgExpect.Environment, func(i, j int) bool {
			return cfgExpect.Environment[i].Name > cfgExpect.Environment[j].Name
		})

		sort.Slice(cfgActual.Environment, func(i, j int) bool {
			return cfgActual.Environment[i].Name > cfgActual.Environment[j].Name
		})

		require.Equal(t, cfgExpect, cfgActual)
	})

	t.Run("ERROR_READING_CONFIG", func(t *testing.T) {
		cfg, err := getFromFile("unreadable config path")
		require.ErrorIs(t, err, os.ErrNotExist)
		require.Equal(t, cfg, NewEmptyConfig())
	})

	t.Run("ERROR_UNMARSHALLING_CONFIG", func(t *testing.T) {
		cfgPath := path.Join(tmpDirPath, path.Base(t.Name())+".yaml")
		defer func() {
			require.NoError(t, os.RemoveAll(cfgPath))
		}()
		require.NoError(t,
			os.WriteFile(
				cfgPath,
				[]byte("1f!cked #p\nc0nfig"),
				os.ModePerm))

		cfg, err := getFromFile(cfgPath)
		require.Equal(t, err.Error(), "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `1f!cked` into matreshka.AppConfig\n\nerror decoding config to struct")
		require.Equal(t, cfg, NewEmptyConfig())
	})
}

func Test_MergeConfigs(t *testing.T) {
	t.Parallel()

	tmpDirPath := path.Join(os.TempDir(), t.Name())
	require.NoError(t, os.MkdirAll(tmpDirPath, os.ModePerm))

	t.Run("OK", func(t *testing.T) {
		// preparing empty config
		emptyConfigPath := path.Join(tmpDirPath, path.Base(t.Name()+"_empty")+".yaml")
		{
			defer func() {
				require.NoError(t, os.RemoveAll(emptyConfigPath))
			}()
			require.NoError(t,
				os.WriteFile(
					emptyConfigPath,
					emptyConfig,
					os.ModePerm))
		}

		fullConfigPath := path.Join(tmpDirPath, path.Base(t.Name()+"_full")+".yaml")
		{
			defer func() {
				require.NoError(t, os.RemoveAll(fullConfigPath))
			}()
			require.NoError(t,
				os.WriteFile(
					fullConfigPath,
					fullConfig,
					os.ModePerm))
		}

		expectedCfg := AppConfig{
			AppInfo: AppInfo{
				Name:            "matreshka",
				Version:         "v0.0.1",
				StartupDuration: time.Second * 10,
			},
			DataSources: []resources.Resource{
				getPostgresClientTest(),
				getRedisClientTest(),
				getTelegramClientTest(),
				getGRPCClientTest(),
			},
			Servers:          getConfigServersFull(),
			Environment:      Environment(getEnvironmentVariables()),
			ServiceDiscovery: getConfigServiceDiscovery(),
		}

		sort.Slice(expectedCfg.Environment, func(i, j int) bool {
			return expectedCfg.Environment[i].Name > expectedCfg.Environment[j].Name
		})

		t.Run("EMPTY_MERGE_FULL", func(t *testing.T) {
			// empty and full config merge
			actualCfg, err := ReadConfigs(emptyConfigPath, fullConfigPath)

			sort.Slice(actualCfg.Environment, func(i, j int) bool {
				return actualCfg.Environment[i].Name > actualCfg.Environment[j].Name
			})

			require.NoError(t, err)
			require.Equal(t, expectedCfg, actualCfg)
		})

		t.Run("FULL_MERGE_EMPTY", func(t *testing.T) {
			actualCfg, err := ReadConfigs(fullConfigPath, emptyConfigPath)

			sort.Slice(actualCfg.Environment, func(i, j int) bool {
				return actualCfg.Environment[i].Name > actualCfg.Environment[j].Name
			})

			require.NoError(t, err)
			require.Equal(t, expectedCfg, actualCfg)
		})
	})

	t.Run("INVALID_READING_ONE_CONFIG", func(t *testing.T) {
		cfgPath := path.Join(tmpDirPath, path.Base(t.Name())+".yaml")

		defer func() {
			require.NoError(t, os.RemoveAll(cfgPath))
		}()

		require.NoError(t,
			os.WriteFile(
				cfgPath,
				emptyConfig,
				os.ModePerm))

		cfg, err := ReadConfigs(cfgPath, "unreadable config path")
		require.ErrorIs(t, err, os.ErrNotExist)
		require.Equal(t, cfg, NewEmptyConfig())
	})

	t.Run("INVALID_READING_FIRST_CONFIG", func(t *testing.T) {
		cfg, err := ReadConfigs("unreadable config path")
		require.ErrorIs(t, err, os.ErrNotExist)
		require.Equal(t, cfg, NewEmptyConfig())
	})
}
