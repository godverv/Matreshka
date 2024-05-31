package resources

const SqliteResourceName = "sqlite"

type Sqlite struct {
	Name `yaml:"resource_name" env:"-"`

	Path string `yaml:"path"`
}

func NewSqlite(n Name) Resource {
	return &Sqlite{
		Name: n,
		Path: "/app/data",
	}
}

func (p *Sqlite) GetType() string {
	return SqliteResourceName
}
