package matreshka

import (
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
	"gopkg.in/yaml.v3"

	"github.com/godverv/matreshka/internal/cases"
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

	names := serverNamer{
		data: map[string]struct{}{},
	}

	for _, port := range ports {
		srv := s[port]

		// TODO Can't handle large words
		srv.Name = names.getName(srv)
		subPrefix := prefix + srv.Name

		if srv.Port == "" {
			srv.Port = strconv.Itoa(port)
		}
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
		serverName := v.Name[len(rootNode.Name)+1:]

		srv := &server.Server{
			Name: serverName,
		}
		err := srv.UnmarshalEnv(v)
		if err != nil {
			return errors.Wrap(err, "error unmarshalling server description")
		}

		p, err := strconv.Atoi(srv.Port)
		if err != nil {
			return errors.Wrap(err, "port must be an integer, got %s", srv.Port)
		}

		s[p] = srv
	}
	return nil
}

func (s Servers) ParseToStruct(dst any) error {
	dstRef := reflect.ValueOf(dst)
	if dstRef.Kind() != reflect.Ptr {
		return errors.Wrap(ErrNotAPointer, "expected destination to be a pointer ")
	}

	dstRef = dstRef.Elem()
	numFields := dstRef.NumField()

	dstMapping := make(map[string]reflect.Value)

	for i := 0; i < numFields; i++ {
		field := dstRef.Type().Field(i)
		dstMapping[strings.ToLower(field.Name)] = dstRef.Field(i)
	}

	for _, serv := range s {
		name := strings.ToLower(ServerName(serv.Name))

		v, ok := dstMapping[name]
		if !ok {
			return errors.New("not found field with name: " + name)
		}

		v.Set(reflect.ValueOf(serv))
	}

	return nil

}

type serverNamer struct {
	data map[string]struct{}
}

func (s *serverNamer) getName(in *server.Server) string {
	var name string
	if in.Name != "" {
		name = in.Name
	} else {
		name = ServerName("")
	}

	idx := 1

	for {
		_, ok := s.data[name]
		if !ok {
			s.data[name] = struct{}{}
			break
		}
		idx++
		name = ServerName("") + strconv.Itoa(idx)
	}

	name = strings.NewReplacer(
		"_", "-",
		" ", "-",
	).Replace(name)
	name = strings.ToUpper(name)
	return name
}

func ServerName(name string) string {
	if name == "" {
		return "Master"
	}

	name = strings.ReplaceAll(name, " ", "_")
	return cases.SnakeToPascal(name)
}
