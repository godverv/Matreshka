package matreshka

import (
	"github.com/Red-Sock/env"

	"github.com/godverv/matreshka/data_sources"
	"github.com/godverv/matreshka/servers"
)

func getPostgresClientTest() *data_sources.Postgres {
	return &data_sources.Postgres{
		Name:    "postgres",
		Host:    "localhost",
		Port:    5432,
		DbName:  "matreshka",
		User:    "matreshka",
		Pwd:     "matreshka",
		SslMode: "disable",
	}
}
func getPostgresClientEnvs() []env.Node {
	pg := getPostgresClientTest()

	prefix := resourcePrefix + pg.GetName()
	return []env.Node{
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

func getRedisClientTest() *data_sources.Redis {
	return &data_sources.Redis{
		Name: "redis",
		Host: "localhost",
		Port: 6379,
		User: "",
		Pwd:  "",
		Db:   0,
	}
}
func getRedisClientEnvs() []env.Node {
	redis := getRedisClientTest()
	name := resourcePrefix + redis.GetName()

	return []env.Node{
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

func getGRPCClientTest() *data_sources.GRPC {
	return &data_sources.GRPC{
		Name:             "grpc_rscli_example",
		ConnectionString: "0.0.0.0:50051",
		Module:           "github.com/Red-Sock/rscli_example",
	}
}
func getGRPCClientEnvs() []env.Node {
	grpcClient := getGRPCClientTest()
	name := resourcePrefix + grpcClient.GetName()
	return []env.Node{
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

func getTelegramClientTest() *data_sources.Telegram {
	return &data_sources.Telegram{
		Name:   "telegram",
		ApiKey: "some_api_key",
	}
}
func getTelegramClientEnvs() []env.Node {
	telegram := getTelegramClientTest()
	name := resourcePrefix + telegram.GetName()
	return []env.Node{
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
func getRestServerEnvs() []env.Node {
	rest := getRestServerTest()
	serverName := apiPrefix + rest.GetName()

	return []env.Node{
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
func getGRPCServerEnvs() []env.Node {
	grpc := getGRPCServerTest()

	serverName := apiPrefix + grpc.GetName()
	return []env.Node{
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
