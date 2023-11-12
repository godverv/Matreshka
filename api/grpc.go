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

func (g *GRPC) ToEnv() map[string]string {
	return map[string]string{
		EnvVarGRPCPort: strconv.FormatUint(uint64(g.Port), 10),
	}
}

func (g *GRPC) FromEnv(in map[string]string) error {
	portUint, err := strconv.ParseUint(in[EnvVarGRPCPort], 10, 16)
	if err != nil {
		return err
	}

	g.Port = uint16(portUint)

	return nil
}

func (g *GRPC) GetPort() uint16 {
	if g.Port != 0 {
		return g.Port
	}

	return DefaultGrpcPort
}

func (g *GRPC) GetPortStr() string {
	p := g.GetPort()
	return strconv.FormatUint(uint64(p), 10)
}
