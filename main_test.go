package matreshka

import (
	_ "embed"

	"github.com/Red-Sock/evon"

	"github.com/godverv/matreshka/resources"
	"github.com/godverv/matreshka/servers"
)

var (
	//go:embed tests/empty_config.yaml
	emptyConfig []byte

	//go:embed tests/resourced_config.yaml
	resourcedConfig []byte
	//go:embed tests/api_config.yaml
	apiConfig []byte
	//go:embed tests/environment_config.yaml
	environmentConfig []byte

	//go:embed tests/full_config.yaml
	fullConfig []byte

	//go:embed tests/.env.full_config
	dotEnvFullConfig []byte
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
		User: "",
		Pwd:  "",
		Db:   0,
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
