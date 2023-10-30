package server

const DefaultRestPort = 8080

type Rest struct {
	Name string `yaml:"name"`
	Port uint16 `yaml:"port"`
}

func (r *Rest) GetName() string {
	return r.Name
}

func (r *Rest) GetPort() uint16 {
	if r.Port != 0 {
		return r.Port
	}

	return DefaultRestPort
}
