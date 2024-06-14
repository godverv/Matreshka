package matreshka

import (
	_ "embed"
	"time"

	"github.com/Red-Sock/evon"

	"github.com/godverv/matreshka/environment"
	"github.com/godverv/matreshka/resources"
	"github.com/godverv/matreshka/servers"
)

var (
	//go:embed tests/empty_config.yaml
	emptyConfig []byte

	//go:embed tests/full_config.yaml
	fullConfig []byte

	//go:embed tests/.env.full_config
	dotEnvFullConfig []byte

	//go:embed tests/go.environment_struct.generated
	goCustomEnvStruct []byte
)

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

func getRestServerTest() *servers.Rest {
	return &servers.Rest{
		Name: "rest",
		Port: 8080,
	}
}
func getRestServerEnvs() []evon.Node {
	rest := getRestServerTest()
	serverName := rest.GetName()

	return []evon.Node{
		{
			Name:  serverName,
			Value: rest,
		},
		{
			Name:  serverName + "_name",
			Value: rest.GetName(),
		},
		{
			Name:  serverName + "_port",
			Value: int(rest.Port),
		},
	}
}

func getGRPCServerTest() *servers.GRPC {
	return &servers.GRPC{
		Name: "grpc",
		Port: 50051,
	}
}
func getGRPCServerEnvs() []evon.Node {
	grpc := getGRPCServerTest()

	serverName := grpc.GetName()
	return []evon.Node{
		{
			Name:  serverName,
			Value: grpc,
		},
		{
			Name:  serverName + "_name",
			Value: grpc.GetName(),
		},
		{
			Name:  serverName + "_port",
			Value: int(grpc.Port),
		},
	}
}

func getEnvironmentVariables() []*environment.Variable {
	return []*environment.Variable{
		{
			Name:  "available ports",
			Type:  environment.VariableTypeInt,
			Value: []int{10, 12, 34, 35, 36, 37, 38, 39, 40},
		},
		{
			Name:  "credit percent",
			Type:  environment.VariableTypeFloat,
			Value: 0.01,
		},
		{
			Name:  "credit percents based on year of birth",
			Type:  environment.VariableTypeFloat,
			Value: []float64{0.01, 0.02, 0.03, 0.04},
		},

		{
			Name:  "database max connections",
			Value: 1,
			Type:  environment.VariableTypeInt,
		},

		{
			Name:  "one of welcome string",
			Type:  environment.VariableTypeStr,
			Value: "one",
			Enum:  []any{"one", "two", "three"},
		},

		{
			Name:  "request timeout",
			Type:  environment.VariableTypeDuration,
			Value: time.Second * 10,
		},

		{
			Name:  "true falser",
			Type:  environment.VariableTypeBool,
			Value: true,
		},

		{
			Name:  "usernames to ban",
			Type:  environment.VariableTypeStr,
			Value: []string{"hacker228", "mothe4acker"},
		},

		{
			Name:  "welcome string",
			Type:  environment.VariableTypeStr,
			Value: "not so basic 🤡 string",
		},
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

			{
				Name: "MATRESHKA_ENVIRONMENT",
				InnerNodes: []*evon.Node{
					{
						Name:  "MATRESHKA_ENVIRONMENT_AVAILABLE-PORTS",
						Value: "[10,12,34,35,36,37,38,39,40]",
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_AVAILABLE-PORTS_TYPE",
						Value: environment.VariableTypeInt,
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_CREDIT-PERCENT",
						Value: "0.01",
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_CREDIT-PERCENTS-BASED-ON-YEAR-OF-BIRTH",
						Value: "[0.01,0.02,0.03,0.04]",
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_CREDIT-PERCENTS-BASED-ON-YEAR-OF-BIRTH_TYPE",
						Value: environment.VariableTypeFloat,
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_CREDIT-PERCENT_TYPE",
						Value: environment.VariableTypeFloat,
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_DATABASE-MAX-CONNECTIONS",
						Value: "1",
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_DATABASE-MAX-CONNECTIONS_TYPE",
						Value: environment.VariableTypeInt,
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_ONE-OF-WELCOME-STRING",
						Value: "one",
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_ONE-OF-WELCOME-STRING_ENUM",
						Value: "[one,two,three]",
					},

					{
						Name:  "MATRESHKA_ENVIRONMENT_ONE-OF-WELCOME-STRING_TYPE",
						Value: environment.VariableTypeStr,
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_REQUEST-TIMEOUT",
						Value: "10s",
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_REQUEST-TIMEOUT_TYPE",
						Value: environment.VariableTypeDuration,
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_TRUE-FALSER",
						Value: "true",
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_TRUE-FALSER_TYPE",
						Value: environment.VariableTypeBool,
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_USERNAMES-TO-BAN",
						Value: "[hacker228,mothe4acker]",
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_USERNAMES-TO-BAN_TYPE",
						Value: environment.VariableTypeStr,
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_WELCOME-STRING",
						Value: "not so basic 🤡 string",
					},
					{
						Name:  "MATRESHKA_ENVIRONMENT_WELCOME-STRING_TYPE",
						Value: environment.VariableTypeStr,
					},
				},
			},
		},
	}
}
