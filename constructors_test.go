package matreshka

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka/api"
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

		cfg, err := ReadConfig(cfgPath)
		require.NoError(t, err)
		require.Equal(t, cfg, NewEmptyConfig())
	})

	t.Run("ERROR_READING_CONFIG", func(t *testing.T) {
		cfg, err := ReadConfig("unreadable config path")
		require.ErrorIs(t, err, os.ErrNotExist)
		require.Nil(t, cfg)
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

		cfg, err := ReadConfig(cfgPath)
		require.Contains(t, err.Error(), "error decoding config to struct\nyaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `1f!cked` into matreshka.AppConfig")
		require.Nil(t, cfg)
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

		fullConfigExpect := &AppConfig{
			AppInfo: AppInfo{
				Name:            "matreshka",
				Version:         "0.0.1",
				StartupDuration: time.Second * 10,
			},
			Resources: []resources.Resource{
				&resources.Postgres{
					Name:    "postgres",
					Host:    "localhost",
					Port:    5432,
					User:    "matreshka",
					Pwd:     "matreshka",
					DbName:  "matreshka",
					SSLMode: "false",
				},
			},
			Servers: []api.Api{
				&api.Rest{
					Name: "rest_server",
					Port: 8080,
				},
			},
			Environment: map[string]interface{}{
				"int":      1,
				"string":   "not so basic 🤡 string",
				"bool":     true,
				"duration": "10s",
			},
		}

		emptyFullCfg, err := ReadConfigs(emptyConfigPath, fullConfigPath)
		require.NoError(t, err)
		require.Equal(t, emptyFullCfg, fullConfigExpect)

		fullEmptyCfg, err := ReadConfigs(fullConfigPath, emptyConfigPath)
		require.NoError(t, err)
		require.Equal(t, fullEmptyCfg, fullConfigExpect)

		require.Equal(t, emptyFullCfg, fullEmptyCfg)
	})

	t.Run("EMPTY_PARAMS", func(t *testing.T) {
		c, e := ReadConfigs()
		require.Nil(t, e)
		require.Nil(t, c)
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
		require.Nil(t, cfg)
	})
}
