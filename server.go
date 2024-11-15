package matreshka

import (
	"sort"
	"strconv"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/server"
)

type Servers map[int]*server.Server

func (s Servers) GetByName(name string) *server.Server {
	for _, serv := range s {
		if serv.Name == name {
			return serv
		}
	}

	return nil
}

func (s Servers) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var portToServers map[string]yaml.Node
	err := unmarshal(&portToServers)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling to yaml.Nodes")
	}

	for portStr, node := range portToServers {
		srv := &server.Server{
			Port: portStr,
		}

		err = node.Decode(&srv)
		if err != nil {
			return errors.Wrap(err, "error decoding server")
		}

		port, err := strconv.Atoi(portStr)
		if err != nil {
			return errors.Wrap(err, "error converting port to int")
		}

		s[port] = srv
	}

	return nil
}

func (s Servers) MarshalEnv(prefix string) ([]*evon.Node, error) {
	root := evon.Node{
		Name: prefix,
	}

	ports := make([]int, 0, len(s))
	for port := range s {
		ports = append(ports, port)
	}
	if len(ports) == 0 {
		return nil, nil
	}

	sort.Ints(ports)

	if prefix != "" {
		prefix += "_"
	}

	for _, port := range ports {
		srv := s[port]
		subPrefix := prefix + strconv.Itoa(port)
		serverNodes, err := srv.MarshalEnv(subPrefix)
		if err != nil {
			return nil, errors.Wrap(err, "error marshalling server")
		}

		root.InnerNodes = append(root.InnerNodes, serverNodes...)

	}
	return []*evon.Node{&root}, nil
}

func (s Servers) UnmarshalEnv(rootNode *evon.Node) error {
	for _, v := range rootNode.InnerNodes {
		port := v.Name[len(rootNode.Name)+1:]

		p, err := strconv.Atoi(port)
		if err != nil {
			return errors.Wrap(err, "expected port value to be integer")
		}
		srv := &server.Server{}
		err = srv.UnmarshalEnv(v)
		if err != nil {
			return errors.Wrap(err, "error unmarshalling server description")
		}

		s[p] = srv
	}
	return nil
}
