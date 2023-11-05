package api

const (
	RestServerType  = "rest"
	DefaultRestPort = 8080
)

type Rest struct {
	Name

	Port uint16 `yaml:"port"`
}

func (r *Rest) ToEnv() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (r *Rest) FromEnv(in map[string]string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Rest) GetPort() uint16 {
	if r.Port != 0 {
		return r.Port
	}

	return DefaultRestPort
}
