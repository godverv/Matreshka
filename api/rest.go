package api

import (
	"strconv"
)

const (
	RestServerType  = "rest"
	DefaultRestPort = 8080

	EnvVarRestPort = "REST_PORT"
)

type Rest struct {
	Name `env:"-"`

	Port uint16 `yaml:"port"`
}

func (r *Rest) ToEnv() map[string]string {
	return map[string]string{
		EnvVarRestPort: strconv.FormatUint(uint64(r.Port), 10),
		EnvServerName:  r.GetName(),
	}
}

func (r *Rest) FromEnv(in map[string]string) error {
	portUint, err := strconv.ParseUint(in[EnvVarRestPort], 10, 16)
	if err != nil {
		return err
	}

	r.Port = uint16(portUint)

	if r.Name == "" {
		r.Name = RestServerType
	}

	return nil
}

func (r *Rest) GetPort() uint16 {
	if r.Port != 0 {
		return r.Port
	}

	return DefaultRestPort
}

func (r *Rest) GetPortStr() string {
	p := r.GetPort()
	return strconv.FormatUint(uint64(p), 10)
}

func (r *Rest) GetType() string {
	return RestServerType
}
