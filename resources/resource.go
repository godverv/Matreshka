package resources

import (
	"strings"

	"github.com/godverv/matreshka/internal/env_parser"
)

const EnvResourceName = "resource_name"

type Resource interface {
	env_parser.EnvParser

	// GetName - returns Name defined in config file
	GetName() string
	GetType() string
}

type Name string

func (a Name) GetName() string {
	return string(a)
}

func GetResourceByName(name string) Resource {
	switch strings.Split(name, "_")[0] {
	case PostgresResourceName:
		return &Postgres{}

	case RedisResourceName:
		return &Redis{}

	case TelegramResourceName:
		return &Telegram{}

	default:
		return &Unknown{}
	}
}
