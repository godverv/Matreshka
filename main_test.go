package matreshka

import (
	"github.com/godverv/matreshka/api"
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
		SSLMode: "disable",
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

func getGRPCClientTest() *resources.GRPC {
	return &resources.GRPC{
		Name:             "grpc_rscli_example",
		ConnectionString: "0.0.0.0:50051",
		Module:           "github.com/Red-Sock/rscli_example",
	}
}

func getTelegramClientTest() *resources.Telegram {
	return &resources.Telegram{
		Name:   "telegram",
		ApiKey: "some_api_key",
	}
}

func getRestServerTest() *api.Rest {
	return &api.Rest{
		Name: "rest_server",
		Port: 8080,
	}
}

func getGRPCServerTest() *api.GRPC {
	return &api.GRPC{
		Name: "grpc_server",
		Port: 50051,
	}
}
