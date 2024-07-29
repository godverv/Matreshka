package matreshka

import (
	"github.com/godverv/matreshka/service_discovery"
)

type ServiceDiscovery struct {
	// MakoshURl
	// MakoshToken
	Overrides service_discovery.Overrides `yaml:"overrides"`
}
