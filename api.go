package matreshka

import (
	"github.com/godverv/matreshka/api"
)

type Servers []api.Api

func (r Servers) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var resourceNames []api.Name
	err := unmarshal(&resourceNames)
	if err != nil {
		return err
	}

	actualResources := make([]api.Api, 0, len(resourceNames))

	for _, appServer := range resourceNames {
		actualResources = append(actualResources, api.GetServerByName(appServer.GetName()))
	}

	return unmarshal(&actualResources)
}
