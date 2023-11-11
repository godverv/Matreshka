package matreshka

import (
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/resources"
)

type Resources []resources.Resource

func (r *Resources) Get(name string) resources.Resource {
	for _, item := range *r {
		if item.GetName() == name {
			return item
		}
	}

	return nil
}

func (r *Resources) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var resourceNodes []yaml.Node
	err := unmarshal(&resourceNodes)
	if err != nil {
		return err
	}

	actualResources := make([]resources.Resource, len(resourceNodes))

	for resIdx, node := range resourceNodes {
		if len(node.Content) == 0 {
			continue
		}

		actualResources[resIdx] = resources.GetResourceByName(findResourceName(node.Content))
		err = node.Decode(actualResources[resIdx])
		if err != nil {
			return err
		}
	}

	*r = actualResources
	return nil
}

func findResourceName(nodes []*yaml.Node) string {
	for dataIdx := 0; dataIdx < len(nodes); dataIdx += 2 {
		if nodes[dataIdx].Value == "resource_name" {
			return nodes[dataIdx+1].Value
		}
	}

	return ""
}
