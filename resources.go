package matreshka

import (
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/resources"
)

type Resources []resources.Resource

func (r *Resources) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var resourceNodes []yaml.Node
	err := unmarshal(&resourceNodes)
	if err != nil {
		return err
	}

	actualResources := make(Resources, len(resourceNodes))

	for resIdx, node := range resourceNodes {
		if len(node.Content) == 0 {
			continue
		}

		var resourceType string

		for dataIdx := 0; dataIdx < len(node.Content); dataIdx += 2 {
			if node.Content[dataIdx].Value == "resource_name" {
				resourceType = node.Content[dataIdx+1].Value
				break
			}
		}

		actualResources[resIdx] = resources.GetResourceByName(resourceType)
		err = node.Decode(actualResources[resIdx])
		if err != nil {
			return err
		}
	}

	*r = actualResources
	return nil
}
