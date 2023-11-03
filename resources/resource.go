package resources

import (
	"github.com/godverv/matreshka/internal/env_parser"
)

type Resource interface {
	env_parser.EnvParser

	// GetName - returns Name defined in config file
	GetName() string
	GetType() string
}

type AppResource struct {
	ResourceName string `yaml:"resource_name"`
}

func (r *AppResource) GetName() string {
	return r.ResourceName
}

var nameToType = map[string]Resource{
	PostgresResourceName: &Postgres{},
	RedisResourceName:    &Redis{},
	TelegramResourceName: &Telegram{},
}

func GetResourceByName(name string) Resource {
	r, ok := nameToType[name]
	if ok {
		return r
	}

	return &Unknown{}
}
