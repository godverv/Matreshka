package resources

const SqliteResourceName = "sqlite"

type Sqlite struct {
	Name `yaml:"resource_name" env:"-"`

	Path             string `yaml:"path"`
	MigrationsFolder string `yaml:"migrations_folder,omitempty"`
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
