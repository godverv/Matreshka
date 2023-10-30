package resources

import (
	"github.com/godverv/matreshka/internal/env_parser"
)

type Resource interface {
	// GetName - returns Name defined in config file
	GetName() string

	env_parser.EnvParser
}
