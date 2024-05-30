package servers

import (
	"strings"
)

const EnvServerName = "server_name"

type Api interface {
	// GetName - return a name of server
	GetName() string
	// GetPort - return port or default port
	GetPort() uint16
	GetPortStr() string

	GetType() string
}
type Name string

func (s Name) GetName() string {
	return string(s)
}

var apis = map[string]func(n Name) Api{
	RestServerType: func(n Name) Api { return &Rest{Name: n} },
	GRPSServerType: func(n Name) Api { return &GRPC{Name: n} },
}

func GetServerByName(name string) Api {
	name = strings.ToLower(name)
	a := apis[strings.Split(name, "_")[0]]

	if a == nil {
		return &Unknown{
			Name: Name(name),
		}
	}

	return a(Name(name))
}
