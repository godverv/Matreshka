package server

const (
	GRPSServerType = "{grpc}"
)

type GRPC struct {
	Module  string `yaml:"module" env:"module"`
	Gateway string `yaml:"gateway" env:"gateway"`
}

func (g *GRPC) GetType() string {
	return GRPSServerType
}
