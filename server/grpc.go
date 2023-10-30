package server

const DefaultGrpcPort = 50051

type GRPC struct {
	Name string `yaml:"name"`
	Port uint16 `yaml:"port"`
}

func (r *GRPC) GetName() string {
	return r.Name
}

func (r *GRPC) GetPort() uint16 {
	if r.Port != 0 {
		return r.Port
	}

	return DefaultGrpcPort
}
