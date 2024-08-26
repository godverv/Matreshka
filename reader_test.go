package matreshka

import (
	"os"
	"path"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/resources"
)

func Test_ReadConfig(t *testing.T) {
	t.Parallel()

	tmpDirPath := path.Join(os.TempDir(), t.Name())
	require.NoError(t, os.MkdirAll(tmpDirPath, os.ModePerm))

	t.Run("OK", func(t *testing.T) {
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

		sort.Slice(cfgExpect.DataSources, func(i, j int) bool {
			return cfgExpect.DataSources[i].GetName() > cfgExpect.DataSources[j].GetName()
		})

		sort.Slice(cfgActual.DataSources, func(i, j int) bool {
			return cfgActual.DataSources[i].GetName() > cfgActual.DataSources[j].GetName()
		})

		require.Equal(t, cfgExpect, cfgActual)
	})

	t.Run("OK_READ_FULL_FROM_ENVIRONMENT", func(t *testing.T) {
		require.NoError(t, setupEnvironmentVariables())

		cfgActual, err := ReadConfigs()
		require.NoError(t, err)

		cfgExpect := getFullConfigTest()

		sort.Slice(cfgExpect.DataSources, func(i, j int) bool {
			return cfgExpect.DataSources[i].GetName() > cfgExpect.DataSources[j].GetName()
		})

		sort.Slice(cfgActual.DataSources, func(i, j int) bool {
			return cfgActual.DataSources[i].GetName() > cfgActual.DataSources[j].GetName()
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
		require.Contains(t, err.Error(), "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `1f!cked` into matreshka.AppConfig\nerror decoding config to struct")
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
			Servers:     getConfigServersFull(),
			Environment: Environment(getEnvironmentVariables()),
		}

		t.Run("EMPTY_MERGE_FULL", func(t *testing.T) {
			// empty and full config merge
			actualCfg, err := ReadConfigs(emptyConfigPath, fullConfigPath)
			require.NoError(t, err)
			require.Equal(t, expectedCfg, actualCfg)
		})

		t.Run("FULL_MERGE_EMPTY", func(t *testing.T) {
			actualCfg, err := ReadConfigs(fullConfigPath, emptyConfigPath)
			require.NoError(t, err)
			require.Equal(t, expectedCfg, actualCfg)
		})

		t.Run("FULL_MERGE_ENV", func(t *testing.T) {
			//require.NoError(t, setupEnvironmentVariables())
			//gotCfg, err := ReadConfigs(fullConfigPath, emptyConfigPath)
			//require.NoError(t, err)
			//require.Equal(t, gotCfg, expectedCfg)
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
