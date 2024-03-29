package api

import (
	"strings"

	"github.com/godverv/matreshka/internal/env_parser"
)

const EnvServerName = "server_name"

type Api interface {
	// GetName - return a name of server
	GetName() string
	// GetPort - return port or default port
	GetPort() uint16
	GetPortStr() string

	env_parser.EnvParser
}
type Name string

func (s Name) GetName() string {
	return string(s)
}

func GetServerByName(name string) Api {
	switch strings.Split(name, "_")[0] {
	case RestServerType:
		return &Rest{
			Name: Name(name),
		}

	case GRPSServerType:
		return &GRPC{
			Name: Name(name),
		}

	default:
		return &Unknown{
			Name: Name(name),
		}
	}
}
