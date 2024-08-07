package resources

const GrpcResourceName = "grpc"

type GRPC struct {
	Name `yaml:"resource_name" env:"-"`

	ConnectionString string `yaml:"connection_string"`
	Module           string `yaml:"module"`
}

func NewGRPC(n Name) Resource {
	return &GRPC{
		Name: n,
	}
}

func (g *GRPC) GetType() string {
	return GrpcResourceName
}

func (g *GRPC) Obfuscate() Resource {
	return &GRPC{
		Name:             g.Name,
		ConnectionString: "localhost:50051",
		Module:           g.Module,
	}
}
