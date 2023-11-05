package api

const (
	GRPSServerType  = "grpc"
	DefaultGrpcPort = 50051
)

type GRPC struct {
	Name

	Port uint16 `yaml:"port"`
}

func (r *GRPC) ToEnv() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (r *GRPC) FromEnv(in map[string]string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *GRPC) GetPort() uint16 {
	if r.Port != 0 {
		return r.Port
	}

	return DefaultGrpcPort
}
