package matreshka

import (
	"strings"

	"github.com/Red-Sock/env"
	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/data_sources"
)

type DataSources []data_sources.Resource

func (r *DataSources) Postgres(name string) (out *data_sources.Postgres, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*data_sources.Postgres)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) Telegram(name string) (out *data_sources.Telegram, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*data_sources.Telegram)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) Redis(name string) (out *data_sources.Redis, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*data_sources.Redis)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) GRPC(name string) (out *data_sources.GRPC, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*data_sources.GRPC)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) Sqlite(name string) (out *data_sources.Sqlite, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*data_sources.Sqlite)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var resourceNodes []yaml.Node
	err := unmarshal(&resourceNodes)
	if err != nil {
		return err
	}

	actualResources := make([]data_sources.Resource, len(resourceNodes))

	for resIdx, node := range resourceNodes {
		if len(node.Content) == 0 {
			continue
		}

		actualResources[resIdx] = data_sources.GetResourceByName(findResourceName(node.Content))
		err = node.Decode(actualResources[resIdx])
		if err != nil {
			return err
		}
	}

	*r = actualResources
	return nil
}

func (r *DataSources) MarshalEnv(prefix string) []env.Node {
	if prefix != "" {
		prefix += "_"
	}

	out := make([]env.Node, 0, len(*r))
	for _, resource := range *r {
		resourceName := strings.Replace(resource.GetName(), "_", "-", -1)
		out = append(out, env.MarshalEnvWithPrefix(prefix+resourceName, resource)...)
	}

	return out
}
func (r *DataSources) UnmarshalEnv(rootNode *env.Node) error {
	sources := make(DataSources, 0)
	for _, dataSourceNode := range rootNode.InnerNodes {
		name := dataSourceNode.Name

		if strings.HasPrefix(dataSourceNode.Name, rootNode.Name) {
			name = name[len(rootNode.Name)+1:]
		}

		name = strings.Replace(name, "-", "_", -1)

		dst := data_sources.GetResourceByName(name)

		env.NodeToStruct(dataSourceNode.Name, dataSourceNode, dst)
		sources = append(sources, dst)
	}

	*r = sources

	return nil
}

func (r *DataSources) get(name string) data_sources.Resource {
	name = strings.TrimPrefix(name, resourcePrefix)
	for _, item := range *r {
		if item.GetName() == name {
			return item
		}
	}

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
