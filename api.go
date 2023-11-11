package matreshka

import (
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/api"
)

type Servers []api.Api

func (s *Servers) Get(name string) api.Api {
	for _, item := range *s {
		if item.GetName() == name {
			return item
		}
	}

	return nil
}

func (s *Servers) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var apiNodes []yaml.Node
	err := unmarshal(&apiNodes)
	if err != nil {
		return err
	}

	actualApi := make([]api.Api, len(apiNodes))

	for apiIdx, apiNode := range apiNodes {
		if len(apiNode.Content) == 0 {
			continue
		}

		var apiName string
		for nodeIdx := 0; nodeIdx < len(apiNode.Content); nodeIdx += 2 {
			if apiNode.Content[nodeIdx].Value == "name" {
				apiName = apiNode.Content[nodeIdx+1].Value
				break
			}
		}

		actualApi[apiIdx] = api.GetServerByName(apiName)
		err = apiNode.Decode(actualApi[apiIdx])
		if err != nil {
			return err
		}
	}

	*s = actualApi

	return nil
}
