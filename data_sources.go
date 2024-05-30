package matreshka

import (
	"strings"

	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/internal/env"
	"github.com/godverv/matreshka/resources"
)

type DataSources []resources.Resource

func (r *DataSources) Postgres(name string) (out *resources.Postgres, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*resources.Postgres)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) Telegram(name string) (out *resources.Telegram, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*resources.Telegram)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) Redis(name string) (out *resources.Redis, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*resources.Redis)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) GRPC(name string) (out *resources.GRPC, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*resources.GRPC)
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedType, "required type %T got %T", out, res)
	}

	return out, nil
}

func (r *DataSources) Sqlite(name string) (out *resources.Sqlite, err error) {
	res := r.get(name)
	if res == nil {
		return nil, ErrNotFound
	}

	out, ok := res.(*resources.Sqlite)
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

func (r *DataSources) MarshalEnv(prefix string) []env.EnvVal {
	if prefix != "" {
		prefix += "_"
	}

	out := make([]env.EnvVal, 0, len(*r))
	for _, resource := range *r {
		out = append(out, env.MarshalEnvWithPrefix(prefix+resource.GetName(), resource)...)
	}

	return out
}
func (r *DataSources) UnmarshalEnv(env *env.EnvNode) error {
	return nil
}
func (r *DataSources) get(name string) resources.Resource {
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
