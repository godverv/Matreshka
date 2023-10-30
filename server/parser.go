package server

import (
	"github.com/godverv/matreshka/internal/env_parser"
)

type Server interface {
	// GetName - return a name of server
	GetName() string
	// GetPort - return port or default port
	GetPort() uint16

	env_parser.EnvParser
}
