package matreshka

import (
	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/internal/env"
	"github.com/godverv/matreshka/resources"
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
func getPostgresClientEnvs() []env.EnvVal {
	pg := getPostgresClientTest()

	prefix := resourcePrefix + pg.GetName()
	return []env.EnvVal{
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
func getRedisClientEnvs() []env.EnvVal {
	redis := getRedisClientTest()
	name := resourcePrefix + redis.GetName()

	return []env.EnvVal{
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
func getGRPCClientEnvs() []env.EnvVal {
	grpcClient := getGRPCClientTest()
	name := resourcePrefix + grpcClient.GetName()
	return []env.EnvVal{
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
func getTelegramClientEnvs() []env.EnvVal {
	telegram := getTelegramClientTest()
	name := resourcePrefix + telegram.GetName()
	return []env.EnvVal{
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

func getRestServerTest() *api.Rest {
	return &api.Rest{
		Name: "rest",
		Port: 8080,
	}
}
func getRestServerEnvs() []env.EnvVal {
	rest := getRestServerTest()
	serverName := apiPrefix + rest.GetName()

	return []env.EnvVal{
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

func getGRPCServerTest() *api.GRPC {
	return &api.GRPC{
		Name: "grpc",
		Port: 50051,
	}
}
func getGRPCServerEnvs() []env.EnvVal {
	grpc := getGRPCServerTest()

	serverName := apiPrefix + grpc.GetName()
	return []env.EnvVal{
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
