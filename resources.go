package matreshka

import (
	"github.com/godverv/matreshka/resources"
)

type Resources []resources.Resource

func (r Resources) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var resourceNames []resources.Name
	err := unmarshal(&resourceNames)
	if err != nil {
		return err
	}

	actualResources := make([]resources.Resource, 0, len(resourceNames))

	for _, appResource := range resourceNames {
		actualResources = append(actualResources, resources.GetResourceByName(appResource.GetName()))
	}

	return unmarshal(&actualResources)
}
