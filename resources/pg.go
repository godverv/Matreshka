package resources

import (
	"fmt"
	"strings"

	"go.redsock.ru/evon"
	"go.redsock.ru/rerrors"
)

const PostgresResourceName = "postgres"

type Postgres struct {
	Name `yaml:"resource_name" env:"-"`

	MigrationsFolder `yaml:"migrations_folder"`

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
	resourceType := strings.Split(p.GetName(), evon.ObjectSplitter)[0]
	if resourceType != "postgres" {
		return nil, rerrors.Wrap(ErrInvalidResourceName, "but got: "+resourceType)
	}

	return *p, nil
}

func (p *Postgres) ConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		p.User,
		p.Pwd,
		p.Host,
		p.Port,
		p.DbName,
	)
}

func (p *Postgres) SqlDialect() string {
	return "postgres"
}
