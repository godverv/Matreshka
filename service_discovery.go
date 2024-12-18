package matreshka

import (
	"go.verv.tech/matreshka/service_discovery"
)

type ServiceDiscovery struct {
	// MakoshURl
	// MakoshToken
	Overrides service_discovery.Overrides `yaml:"overrides"`
}
