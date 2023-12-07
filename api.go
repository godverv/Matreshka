package matreshka

import (
	"strings"

	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/api"
)

type Servers []api.Api

func (s *Servers) GRPC(name string) (*api.GRPC, error) {
	res := s.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*api.GRPC)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (s *Servers) REST(name string) (*api.Rest, error) {
	res := s.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*api.Rest)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (s *Servers) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var apiNodes []yaml.Node
	err := unmarshal(&apiNodes)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling to yaml.Nodes")
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
			return errors.Wrapf(err, "error decoding to struct of type %T", actualApi[apiIdx])
		}
	}

	*s = actualApi

	return nil
}

func (s *Servers) get(name string) api.Api {
	name = strings.TrimLeft(name, apiPrefix)
	for _, item := range *s {
		if item.GetName() == name {
			return item
		}
	}

	return nil
}
