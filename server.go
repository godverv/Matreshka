package matreshka

import (
	"strings"

	"github.com/Red-Sock/env"
	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/servers"
)

type Servers []servers.Api

func (s *Servers) GRPC(name string) (*servers.GRPC, error) {
	res := s.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*servers.GRPC)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (s *Servers) REST(name string) (*servers.Rest, error) {
	res := s.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*servers.Rest)
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

	actualApi := make([]servers.Api, len(apiNodes))

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

		actualApi[apiIdx] = servers.GetServerByName(apiName)
		err = apiNode.Decode(actualApi[apiIdx])
		if err != nil {
			return errors.Wrapf(err, "error decoding to struct of type %T", actualApi[apiIdx])
		}
	}

	*s = actualApi

	return nil
}

func (s *Servers) MarshalEnv(prefix string) []env.Node {
	if prefix != "" {
		prefix += "_"
	}

	out := make([]env.Node, 0)
	for _, srv := range *s {
		serverName := strings.Replace(srv.GetName(), "_", "-", -1)
		out = append(out, env.MarshalEnvWithPrefix(prefix+serverName, srv)...)
	}

	return out
}
func (s *Servers) UnmarshalEnv(rootNode *env.Node) error {
	srvs := make(Servers, 0)
	for _, serverNode := range rootNode.InnerNodes {
		name := serverNode.Name

		if strings.HasPrefix(serverNode.Name, rootNode.Name) {
			name = name[len(rootNode.Name)+1:]
		}

		name = strings.Replace(name, "-", "_", -1)

		dst := servers.GetServerByName(name)

		env.NodeToStruct(serverNode.Name, serverNode, dst)
		srvs = append(srvs, dst)
	}

	*s = srvs

	return nil
}

func (s *Servers) get(name string) servers.Api {
	name = strings.TrimLeft(name, apiPrefix)
	for _, item := range *s {
		if item.GetName() == name {
			return item
		}
	}

	return nil
}
