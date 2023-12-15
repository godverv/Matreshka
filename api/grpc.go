package api

import (
	"strconv"
)

const (
	GRPSServerType  = "grpc"
	DefaultGrpcPort = 50051

	EnvVarGRPCPort        = "GRPC_PORT"
	EnvVarGRPCGatewayPort = "GRPC_GATEWAY_PORT"
)

type GRPC struct {
	Name

	Port        uint16 `yaml:"port"`
	GatewayPort uint16 `yaml:"gateway_port"`
}

func (g *GRPC) ToEnv() map[string]string {
	return map[string]string{
		EnvVarGRPCPort:        strconv.FormatUint(uint64(g.Port), 10),
		EnvVarGRPCGatewayPort: strconv.FormatUint(uint64(g.GatewayPort), 10),
		EnvServerName:         g.GetName(),
	}
}

func (g *GRPC) FromEnv(in map[string]string) error {
	portUint, err := strconv.ParseUint(in[EnvVarGRPCPort], 10, 16)
	if err != nil {
		return err
	}

	g.Port = uint16(portUint)

	gwPortUint, err := strconv.ParseUint(in[EnvVarGRPCGatewayPort], 10, 16)
	if err != nil {
		return err
	}

	g.GatewayPort = uint16(gwPortUint)

	if g.Name == "" {
		g.Name = GRPSServerType
	}

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

func (g *GRPC) GetGatewayPort() uint16 {
	return g.GatewayPort
}

func (g *GRPC) GetGatewayPortStr() string {
	p := g.GatewayPort
	return strconv.FormatUint(uint64(p), 10)
}
