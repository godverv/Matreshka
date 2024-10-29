package resources

const SqliteResourceName = "sqlite"

type Sqlite struct {
	Name `yaml:"resource_name" env:"-"`

	Path             string `yaml:"path"`
	MigrationsFolder `yaml:"migrations_folder,omitempty"`
}

func NewSqlite(n Name) Resource {
	return &Sqlite{
		Name:             n,
		Path:             "/app/data",
		MigrationsFolder: "./migrations",
	}
}

func (p *Sqlite) GetType() string {
	return SqliteResourceName
}

func (p *Sqlite) ConnectionString() string {
	return p.Path
}

func (p *Sqlite) SqlDialect() string {
	return "sqlite"
}
