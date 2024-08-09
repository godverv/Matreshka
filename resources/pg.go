package resources

import (
	"strings"

	errors "github.com/Red-Sock/trace-errors"
)

const PostgresResourceName = "postgres"

type Postgres struct {
	Name `yaml:"resource_name" env:"-"`

	Host string `yaml:"host"`
	Port uint64 `yaml:"port"`

	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`

	DbName  string `yaml:"name"`
	SslMode string `yaml:"ssl_mode"`
}

func NewPostgres(n Name) Resource {
	return &Postgres{
		Name:   n,
		Host:   "0.0.0.0",
		Port:   5432,
		User:   "postgres",
		Pwd:    "",
		DbName: "postgres",
	}
}

func (p *Postgres) GetType() string {
	return PostgresResourceName
}

func (p *Postgres) MarshalYAML() (interface{}, error) {
	resourceType := strings.Split(p.GetName(), "_")[0]
	if resourceType != "postgres" {
		return nil, errors.Wrap(ErrInvalidResourceName, "but got: "+resourceType)
	}

	return *p, nil
}

func (p *Postgres) Obfuscate() Resource {
	return &Postgres{
		Name:    p.Name,
		Host:    "localhost",
		Port:    5432,
		User:    "postgres",
		Pwd:     "postgres",
		DbName:  "master",
		SslMode: "",
	}
}
