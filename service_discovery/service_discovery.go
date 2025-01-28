package service_discovery

import (
	"strings"

	"go.redsock.ru/evon"
	"go.redsock.ru/rerrors"
)

type Override struct {
	ServiceName string `yaml:"service_name"`
	Urls        Urls   `yaml:"urls"`
}

type Overrides []*Override

type Urls []string

func (o *Overrides) UnmarshalEnv(rootNode *evon.Node) error {
	overrides := make(Overrides, 0)
	for _, overrideNode := range rootNode.InnerNodes {
		name := overrideNode.Name

		if strings.HasPrefix(overrideNode.Name, rootNode.Name) {
			name = name[len(rootNode.Name)+1:]
		}

		name = strings.Replace(name, evon.FieldSplitter, evon.ObjectSplitter, -1)

		dst := &Override{}

		err := evon.NodeToStruct(overrideNode.Name, overrideNode, dst)
		if err != nil {
			return rerrors.Wrap(err, "error unmarshalling resource from env")
		}
		overrides = append(overrides, dst)
	}

	*o = overrides

	return nil
}
func (o Overrides) MarshalEnv(prefix string) ([]*evon.Node, error) {
	if prefix != "" {
		prefix += evon.ObjectSplitter
	}

	out := make([]*evon.Node, 0, len(o))
	for _, override := range o {
		overrideServiceName := strings.Replace(override.ServiceName, evon.ObjectSplitter, evon.FieldSplitter, -1)

		nodes, err := evon.MarshalEnvWithPrefix(prefix+overrideServiceName, override)
		if err != nil {
			return nil, rerrors.Wrap(err, "error marshalling service discovery override")
		}

		out = append(out, nodes)
	}

	return out, nil
}

func (o *Override) UnmarshalEnv(n *evon.Node) error {
	return nil
}
func (o *Override) MarshalEnv(prefix string) ([]*evon.Node, error) {
	return nil, nil
}

func (u *Urls) UnmarshalEnv(n *evon.Node) error {
	switch v := n.Value.(type) {
	case string:
		*u = strings.Split(v, " ")
		return nil
	default:
		return rerrors.New("not a string value")
	}

	return nil
}
