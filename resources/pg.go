package resources

import (
	"strconv"
	"strings"

	errors "github.com/Red-Sock/trace-errors"
)

const PostgresResourceName = "postgres"

const (
	EnvVarPostgresHost     = "POSTGRES_HOST"
	EnvVarPostgresPort     = "POSTGRES_PORT"
	EnvVarPostgresUser     = "POSTGRES_USER"
	EnvVarPostgresPassword = "POSTGRES_PWD"
	EnvVarPostgresDbName   = "POSTGRES_NAME"
)

type Postgres struct {
	Name `yaml:"resource_name"`

	Host string `yaml:"host"`
	Port uint64 `yaml:"port"`

	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`

	DbName  string `yaml:"name"`
	SSLMode string `yaml:"ssl_mode"`
}

func (p *Postgres) GetType() string {
	return PostgresResourceName
}

func (p *Postgres) ToEnv() map[string]string {
	return map[string]string{
		EnvVarPostgresUser:     p.User,
		EnvVarPostgresPassword: p.Pwd,
		EnvVarPostgresDbName:   p.DbName,

		EnvVarPostgresHost: p.Host,
		EnvVarPostgresPort: strconv.FormatUint(p.Port, 10),
	}
}

func (p *Postgres) FromEnv(in map[string]string) (err error) {
	p.User = in[EnvVarPostgresUser]
	p.Pwd = in[EnvVarPostgresPassword]
	p.DbName = in[EnvVarPostgresDbName]

	p.Host = in[EnvVarPostgresHost]
	p.Port, err = strconv.ParseUint(in[EnvVarPostgresPort], 10, 64)
	if err != nil {
		return errors.Wrap(err, "error parsing port value")
	}
	return nil
}

func (p *Postgres) MarshalYAML() (interface{}, error) {
	resourceType := strings.Split(p.GetName(), "_")[0]
	if resourceType != "postgres" {
		return nil, errors.Wrap(ErrInvalidResourceName, "but got: "+resourceType)
	}

	return *p, nil
}
