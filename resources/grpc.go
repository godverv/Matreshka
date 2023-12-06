package resources

const GrpcResourceName = "grpc"

const (
	EnvGrpcConnectionString = `GRPC_CONNECTION_STRING`
	EnvGrpcModule           = `GRPC_PACKAGE`
)

type GRPC struct {
	Name `yaml:"resource_name"`

	ConnectionString string `yaml:"connection_string"`
	Module           string `yaml:"module"`
}

func (g *GRPC) GetType() string {
	return GrpcResourceName
}

func (g *GRPC) ToEnv() map[string]string {
	return map[string]string{
		EnvResourceName: g.GetName(),

		EnvGrpcConnectionString: g.ConnectionString,
		EnvGrpcModule:           g.Module,
	}
}

func (g *GRPC) FromEnv(in map[string]string) (err error) {
	g.Name = Name(in[EnvResourceName])

	g.ConnectionString = in[EnvGrpcConnectionString]

	return nil
}
