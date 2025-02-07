package matreshka

import (
	"bytes"
	_ "embed"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.redsock.ru/evon"

	"go.vervstack.ru/matreshka/environment"
	"go.vervstack.ru/matreshka/resources"
)

var (
	//go:embed tests/empty_config.yaml
	emptyConfig []byte

	//go:embed tests/app_info_config_short_name.yaml
	appInfoConfigShortName []byte
	//go:embed tests/app_info_config_full_name.yaml
	appInfoConfigFullName []byte

	//go:embed tests/api_config.yaml
	apiConfig []byte
	//go:embed tests/api_config.with_name.yaml
	apiConfigWithName []byte
	//go:embed tests/.env.api_config
	apiEnvConfig []byte
	//go:embed tests/api_config.invalid_port.yaml
	apiInvalidPortConfig []byte
	//go:embed tests/api_config.invalid.yaml
	apiInvalidStructConfig []byte
	//go:embed tests/api_config.invalid_item.yaml
	apiInvalidItemConfig []byte

	//go:embed tests/full_config.yaml
	fullConfig []byte
	//go:embed tests/full_config_without_module_name.env
	fullEnvConfigWithoutModule []byte
	//go:embed tests/full_config.env
	fullEnvConfig []byte

	//go:embed tests/service_discovery.yaml
	serviceDiscoveryConfig []byte
	//go:embed tests/service_discovery.env
	serviceDiscoveryEnvConfig []byte

	//go:embed tests/full_resources_config.yaml
	fullResourcesConfig []byte
	//go:embed tests/full_resources_config.env
	fullResourcesEnvConfig []byte
)

func getResourcesFull() []resources.Resource {
	return []resources.Resource{
		getPostgresClientTest(),
		getRedisClientTest(),
		getTelegramClientTest(),
		getGRPCClientTest(),
	}
}

func getPostgresClientTest() *resources.Postgres {
	return &resources.Postgres{
		Name:    "postgres",
		Host:    "localhost",
		Port:    5432,
		DbName:  "matreshka",
		User:    "matreshka",
		Pwd:     "matreshka",
		SslMode: "disable",
	}
}
func getPostgresClientEnvs() []evon.Node {
	pg := getPostgresClientTest()

	prefix := pg.GetName()
	return []evon.Node{
		{
			Name:  prefix,
			Value: pg,
		},
		{
			Name:  prefix + "_resource_name",
			Value: pg.GetName(),
		},
		{
			Name:  prefix + "_host",
			Value: pg.Host,
		},
		{
			Name:  prefix + "_port",
			Value: int(pg.Port),
		},
		{
			Name:  prefix + "_user",
			Value: pg.User,
		},
		{
			Name:  prefix + "_pwd",
			Value: pg.Pwd,
		},
		{
			Name:  prefix + "_name",
			Value: pg.DbName,
		},
		{
			Name:  prefix + "_ssl_mode",
			Value: pg.SslMode,
		},
	}
}

func getRedisClientTest() *resources.Redis {
	return &resources.Redis{
		Name: "redis",
		Host: "localhost",
		Port: 6379,
		User: "redis_matreshka",
		Pwd:  "redis_matreshka_pwd",
		Db:   2,
	}
}
func getRedisClientEnvs() []evon.Node {
	redis := getRedisClientTest()
	name := redis.GetName()

	return []evon.Node{
		{
			Name:  name,
			Value: redis,
		},
		{
			Name:  name + "_user",
			Value: redis.User,
		},
		{
			Name:  name + "_resource_name",
			Value: redis.GetName(),
		},
		{
			Name:  name + "_pwd",
			Value: redis.Pwd,
		},
		{
			Name:  name + "_host",
			Value: redis.Host,
		},
		{
			Name:  name + "_port",
			Value: int(redis.Port),
		},
		{
			Name:  name + "_db",
			Value: redis.Db,
		},
	}
}

func getGRPCClientTest() *resources.GRPC {
	return &resources.GRPC{
		Name:             "grpc_rscli_example",
		ConnectionString: "0.0.0.0:50051",
		Module:           "github.com/Red-Sock/rscli_example",
	}
}
func getGRPCClientEnvs() []evon.Node {
	grpcClient := getGRPCClientTest()
	name := grpcClient.GetName()
	return []evon.Node{
		{
			Name:  name,
			Value: grpcClient,
		},
		{
			Name:  name + "_connection_string",
			Value: grpcClient.ConnectionString,
		},
		{
			Name:  name + "_module",
			Value: grpcClient.Module,
		},
		{
			Name:  name + "_resource_name",
			Value: grpcClient.GetName(),
		},
	}
}

func getTelegramClientTest() *resources.Telegram {
	return &resources.Telegram{
		Name:   "telegram",
		ApiKey: "some_api_key",
	}
}
func getTelegramClientEnvs() []evon.Node {
	telegram := getTelegramClientTest()
	name := telegram.GetName()
	return []evon.Node{
		{
			Name:  name,
			Value: telegram,
		},
		{
			Name:  name + "_api_key",
			Value: telegram.ApiKey,
		},

		{
			Name:  name + "_resource_name",
			Value: telegram.GetName(),
		},
	}
}

func getEnvironmentVariables() []*environment.Variable {
	return []*environment.Variable{
		environment.MustNewVariable("database_max_connections", 1),
		environment.MustNewVariable("welcome_string", "not so basic 🤡 string"),
		environment.MustNewVariable("one_of_welcome_string", "one", environment.WithEnum("one", "two", "three")),
		environment.MustNewVariable("true_falser", true),
		environment.MustNewVariable("request_timeout", time.Second*10),
		environment.MustNewVariable("available_ports", []int{10, 12, 34, 35, 36, 37, 38, 39, 40}),
		environment.MustNewVariable("usernames_to_ban", []string{"hacker228", "mothe4acker"}),
		environment.MustNewVariable("credit_percent", 0.01),
		environment.MustNewVariable("credit_percents_based_on_year_of_birth", []float64{0.01, 0.02, 0.03, 0.04}),
	}
}

func getEvonFullConfig() *evon.Node {
	return &evon.Node{
		Name: "MATRESHKA",
		InnerNodes: []*evon.Node{
			// APP INFO
			{
				Name: "MATRESHKA_APP-INFO",
				InnerNodes: []*evon.Node{
					{
						Name:  "MATRESHKA_APP-INFO_NAME",
						Value: "matreshka",
					},
					{
						Name:  "MATRESHKA_APP-INFO_VERSION",
						Value: "v0.0.1",
					},
					{
						Name:  "MATRESHKA_APP-INFO_STARTUP-DURATION",
						Value: time.Second * 10,
					},
				},
			},
			// Data sources
			{
				Name: "MATRESHKA_DATA-SOURCES",
				InnerNodes: []*evon.Node{
					{
						Name: "MATRESHKA_DATA-SOURCES_POSTGRES",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_DATA-SOURCES_POSTGRES_HOST",
								Value: "localhost",
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_POSTGRES_PORT",
								Value: uint64(5432),
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_POSTGRES_USER",
								Value: "matreshka",
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_POSTGRES_PWD",
								Value: "matreshka",
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_POSTGRES_DB-NAME",
								Value: "matreshka",
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_POSTGRES_SSL-MODE",
								Value: "disable",
							},
						},
					},
					{
						Name: "MATRESHKA_DATA-SOURCES_REDIS",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_DATA-SOURCES_REDIS_HOST",
								Value: "localhost",
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_REDIS_PORT",
								Value: uint16(6379),
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_REDIS_USER",
								Value: "redis_matreshka",
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_REDIS_PWD",
								Value: "redis_matreshka_pwd",
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_REDIS_DB",
								Value: 2,
							},
						},
					},
					{
						Name: "MATRESHKA_DATA-SOURCES_TELEGRAM",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_DATA-SOURCES_TELEGRAM_API-KEY",
								Value: "some_api_key",
							},
						},
					},
					{
						Name: "MATRESHKA_DATA-SOURCES_GRPC-RSCLI-EXAMPLE",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_DATA-SOURCES_GRPC-RSCLI-EXAMPLE_CONNECTION-STRING",
								Value: "0.0.0.0:50051",
							},
							{
								Name:  "MATRESHKA_DATA-SOURCES_GRPC-RSCLI-EXAMPLE_MODULE",
								Value: "github.com/Red-Sock/rscli_example",
							},
						},
					},
				},
			},
			// SERVERS
			{
				Name: "MATRESHKA_SERVERS",
				InnerNodes: []*evon.Node{
					{
						Name: "MATRESHKA_SERVERS_REST",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_SERVERS_REST_PORT",
								Value: uint16(8080),
							},
						},
					},
					{
						Name: "MATRESHKA_SERVERS_GRPC",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_SERVERS_GRPC_PORT",
								Value: uint16(50051),
							},
						},
					},
				},
			},
			// Environment
			{
				Name: "MATRESHKA_ENVIRONMENT",
				InnerNodes: []*evon.Node{
					{
						Name:  "MATRESHKA_ENVIRONMENT_AVAILABLE-PORTS",
						Value: "[10,12,34,35,36,37,38,39,40]",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_AVAILABLE-PORTS_TYPE",
								Value: environment.VariableTypeInt,
							},
						},
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_CREDIT-PERCENT",
						Value: "0.01",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_CREDIT-PERCENT_TYPE",
								Value: environment.VariableTypeFloat,
							},
						},
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_CREDIT-PERCENTS-BASED-ON-YEAR-OF-BIRTH",
						Value: "[0.01,0.02,0.03,0.04]",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_CREDIT-PERCENTS-BASED-ON-YEAR-OF-BIRTH_TYPE",
								Value: environment.VariableTypeFloat,
							},
						},
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_DATABASE-MAX-CONNECTIONS",
						Value: "1",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_DATABASE-MAX-CONNECTIONS_TYPE",
								Value: environment.VariableTypeInt,
							},
						},
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_ONE-OF-WELCOME-STRING",
						Value: "one",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_ONE-OF-WELCOME-STRING_TYPE",
								Value: environment.VariableTypeStr,
							},
							{
								Name:  "MATRESHKA_ENVIRONMENT_ONE-OF-WELCOME-STRING_ENUM",
								Value: "[one,two,three]",
							},
						},
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_REQUEST-TIMEOUT",
						Value: "10s",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_REQUEST-TIMEOUT_TYPE",
								Value: environment.VariableTypeDuration,
							},
						},
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_TRUE-FALSER",
						Value: "true",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_TRUE-FALSER_TYPE",
								Value: environment.VariableTypeBool,
							},
						},
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_USERNAMES-TO-BAN",
						Value: "[hacker228,mothe4acker]",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_USERNAMES-TO-BAN_TYPE",
								Value: environment.VariableTypeStr,
							},
						},
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_WELCOME-STRING",
						Value: "not so basic 🤡 string",
						InnerNodes: []*evon.Node{
							{
								Name:  "MATRESHKA_ENVIRONMENT_WELCOME-STRING_TYPE",
								Value: environment.VariableTypeStr,
							},
						},
					},
				},
			},
			// ServiceDiscovery
			{
				Name:       "MATRESHKA_SERVICE-DISCOVERY",
				InnerNodes: make([]*evon.Node, 0),
			},
		},
	}
}

func getFullConfigTest() AppConfig {
	cfgExpect := NewEmptyConfig()
	cfgExpect.AppInfo = AppInfo{
		Name:            "matreshka",
		Version:         "v0.0.1",
		StartupDuration: 10 * time.Second,
	}

	cfgExpect.DataSources = getResourcesFull()

	cfgExpect.Servers = getConfigServersFull()

	cfgExpect.Environment = getEnvironmentVariables()

	cfgExpect.ServiceDiscovery = getConfigServiceDiscovery()
	return cfgExpect
}

func setupFullEnvConfigWithModuleName(t *testing.T) {
	if os.Getenv(VervName) != "" {
		return
	}

	err := os.Setenv(VervName, "MATRESHKA")
	require.NoError(t, err)

	setEnvConfigFromFile(t, fullEnvConfig)
}

func setupFullEnvConfigWithoutModuleName(t *testing.T) {
	setEnvConfigFromFile(t, fullEnvConfigWithoutModule)
}

func setEnvConfigFromFile(t *testing.T, cfg []byte) {
	splitedEnvs := bytes.Split(cfg, []byte{'\n'})

	for _, env := range splitedEnvs {
		if len(env) == 0 {
			continue
		}

		nameVal := bytes.Split(env, []byte{'='})
		err := os.Setenv(string(nameVal[0]), string(nameVal[1]))
		require.NoError(t, err)
	}
}
