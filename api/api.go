package api

import (
	"strings"

	"github.com/godverv/matreshka/internal/env_parser"
)

type Api interface {
	// GetName - return a name of server
	GetName() string
	// GetPort - return port or default port
	GetPort() uint16

	env_parser.EnvParser
}
type Name string

func (s Name) GetName() string {
	return string(s)
}

func GetServerByName(name string) Api {
	switch strings.Split(name, "_")[0] {
	case RestServerType:
		return &Rest{}

	case GRPSServerType:
		return &GRPC{}

	default:
		return &Unknown{}
	}
}
