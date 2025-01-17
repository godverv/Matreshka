package matreshka

import (
	"go.verv.tech/matreshka/service_discovery"
)

type ServiceDiscovery struct {
	MakoshUrl   string                      `yaml:"makosh_url" env:",omitempty"`
	MakoshToken string                      `yaml:"makosh_token" env:",omitempty"`
	Overrides   service_discovery.Overrides `yaml:"overrides" env:",omitempty"`
}
