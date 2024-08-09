package resources

import (
	"strings"
)

type Resource interface {
	// GetName - returns Name defined in config file
	GetName() string
	GetType() string
	Obfuscate() Resource
}

type Name string

func (a Name) GetName() string {
	return string(a)
}

var resources = map[string]func(name Name) Resource{
	PostgresResourceName: NewPostgres,
	RedisResourceName:    NewRedis,
	SqliteResourceName:   NewSqlite,

	TelegramResourceName: NewTelegram,
	GrpcResourceName:     NewGRPC,
}

func GetResourceByName(name string) Resource {

	name = strings.ToLower(name)
	resourceTypeName := strings.Split(name, "_")[0]

	r := resources[resourceTypeName]
	if r == nil {
		return &Unknown{
			Name: Name(name),
		}
	}

	return r(Name(name))
}
