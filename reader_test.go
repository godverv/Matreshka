package matreshka

import (
	"os"
	"path"
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

		cfgGot, err := ParseConfig(fullConfig)
		require.NoError(t, err)

		cfgExpect := getFullConfigTest()

		require.Equal(t, cfgExpect, cfgGot)
	})

	t.Run("OK_READ_FULL_FROM_ENVIRONMENT", func(t *testing.T) {
		require.NoError(t, setupEnvironmentVariables())

		cfgGot, err := ReadConfigs()
		require.NoError(t, err)

		cfgExpect := getFullConfigTest()

		require.Equal(t, cfgGot, cfgExpect)
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

		fullConfigExpect := AppConfig{
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
			Environment: Environment(getEnvironmentVariables()),
		}

		t.Run("EMPTY_MERGE_FULL", func(t *testing.T) {
			// empty and full config merge
			gotCfg, err := ReadConfigs(emptyConfigPath, fullConfigPath)
			require.NoError(t, err)
			require.Equal(t, gotCfg, fullConfigExpect)
		})

		t.Run("FULL_MERGE_EMPTY", func(t *testing.T) {
			gotCfg, err := ReadConfigs(fullConfigPath, emptyConfigPath)
			require.NoError(t, err)
			require.Equal(t, gotCfg, fullConfigExpect)
		})

		t.Run("FULL_MERGE_ENV", func(t *testing.T) {
			//require.NoError(t, setupEnvironmentVariables())
			//gotCfg, err := ReadConfigs(fullConfigPath, emptyConfigPath)
			//require.NoError(t, err)
			//require.Equal(t, gotCfg, fullConfigExpect)
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
