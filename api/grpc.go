package api

import (
	"strconv"
)

const (
	GRPSServerType  = "grpc"
	DefaultGrpcPort = 50051

	EnvVarGRPCPort = "GRPC_PORT"
)

type GRPC struct {
	Name

	Port uint16 `yaml:"port"`
}

func (r *GRPC) ToEnv() map[string]string {
	return map[string]string{
		EnvVarGRPCPort: strconv.FormatUint(uint64(r.Port), 10),
	}
}

func (r *GRPC) FromEnv(in map[string]string) error {
	portUint, err := strconv.ParseUint(in[EnvVarGRPCPort], 10, 16)
	if err != nil {
		return err
	}

	r.Port = uint16(portUint)

	return nil
}

func (r *GRPC) GetPort() uint16 {
	if r.Port != 0 {
		return r.Port
	}

	return DefaultGrpcPort
}
